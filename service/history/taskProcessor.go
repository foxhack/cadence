// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package history

import (
	"sync"
	"time"

	workflow "github.com/uber/cadence/.gen/go/shared"
	"github.com/uber/cadence/common"
	"github.com/uber/cadence/common/backoff"
	"github.com/uber/cadence/common/cache"
	"github.com/uber/cadence/common/clock"
	"github.com/uber/cadence/common/log"
	"github.com/uber/cadence/common/log/tag"
	"github.com/uber/cadence/common/metrics"
	"github.com/uber/cadence/common/persistence"
	"github.com/uber/cadence/service/history/config"
	"github.com/uber/cadence/service/history/execution"
	"github.com/uber/cadence/service/history/shard"
	"github.com/uber/cadence/service/history/task"
)

type (
	taskProcessorOptions struct {
		queueSize   int
		workerCount int
	}

	taskInfo struct {
		// TODO: change to queueTaskExecutor
		processor taskExecutor
		task      task.Info

		attempt   int
		startTime time.Time
		logger    log.Logger

		// used by 2DC task life cycle
		// TODO remove when NDC task life cycle is implemented
		shouldProcessTask bool
	}

	taskProcessor struct {
		shard                   shard.Context
		cache                   *execution.Cache
		shutdownCh              chan struct{}
		tasksCh                 chan *taskInfo
		config                  *config.Config
		logger                  log.Logger
		metricsClient           metrics.Client
		timeSource              clock.TimeSource
		retryPolicy             backoff.RetryPolicy
		workerWG                sync.WaitGroup
		domainMetricsScopeCache cache.DomainMetricsScopeCache
		// worker coroutines notification
		workerNotificationChans []chan struct{}
		// duplicate numOfWorker from config.TimerTaskWorkerCount for dynamic config works correctly
		numOfWorker int
		// emit domain tag
		emitMetricsWithDomainTag bool
	}
)

func newTaskInfo(
	processor taskExecutor,
	task task.Info,
	logger log.Logger,
	startTime time.Time,
) *taskInfo {
	return &taskInfo{
		processor:         processor,
		task:              task,
		attempt:           0,
		startTime:         startTime, // used for metrics
		logger:            logger,
		shouldProcessTask: true,
	}
}

func newTaskProcessor(
	options taskProcessorOptions,
	shard shard.Context,
	executionCache *execution.Cache,
	logger log.Logger,
	emitMetricsWithDomainTag bool,
) *taskProcessor {

	workerNotificationChans := []chan struct{}{}
	for index := 0; index < options.workerCount; index++ {
		workerNotificationChans = append(workerNotificationChans, make(chan struct{}, 1))
	}

	base := &taskProcessor{
		shard:                   shard,
		cache:                   executionCache,
		shutdownCh:              make(chan struct{}),
		tasksCh:                 make(chan *taskInfo, options.queueSize),
		config:                  shard.GetConfig(),
		logger:                  logger,
		metricsClient:           shard.GetMetricsClient(),
		domainMetricsScopeCache: shard.GetService().GetDomainMetricsScopeCache(),
		timeSource:              shard.GetTimeSource(),
		workerNotificationChans: workerNotificationChans,
		retryPolicy:             backoff.NewTwoPhaseRetryPolicy(),
		numOfWorker:             options.workerCount,
		emitMetricsWithDomainTag: emitMetricsWithDomainTag,
	}

	return base
}

func (t *taskProcessor) start() {
	for i := 0; i < t.numOfWorker; i++ {
		t.workerWG.Add(1)
		notificationChan := t.workerNotificationChans[i]
		go t.taskWorker(notificationChan)
	}
	t.logger.Info("Task processor started.")
}

func (t *taskProcessor) stop() {
	close(t.shutdownCh)
	if success := common.AwaitWaitGroup(&t.workerWG, time.Minute); !success {
		t.logger.Warn("Task processor timed out on shutdown.")
	}
	t.logger.Info("Task processor shutdown.")
}

func (t *taskProcessor) taskWorker(
	notificationChan chan struct{},
) {
	defer t.workerWG.Done()

	for {
		select {
		case <-t.shutdownCh:
			return
		case task, ok := <-t.tasksCh:
			if !ok {
				return
			}
			t.processTaskAndAck(notificationChan, task)
		}
	}
}

func (t *taskProcessor) retryTasks() {
	for _, workerNotificationChan := range t.workerNotificationChans {
		select {
		case workerNotificationChan <- struct{}{}:
		default:
		}
	}
}

func (t *taskProcessor) addTask(
	task *taskInfo,
) bool {
	// We have a timer to fire.
	select {
	case t.tasksCh <- task:
	case <-t.shutdownCh:
		return false
	}
	return true
}

func (t *taskProcessor) processTaskAndAck(
	notificationChan <-chan struct{},
	task *taskInfo,
) {

	var scope metrics.Scope
	var err error

FilterLoop:
	for {
		select {
		case <-t.shutdownCh:
			// this must return without ack
			return
		default:
			task.shouldProcessTask, err = task.processor.getTaskFilter()(task.task)
			if err == nil {
				break FilterLoop
			}
			time.Sleep(loadDomainEntryForTimerTaskRetryDelay)
		}
	}

	op := func() error {
		scope, err = t.processTaskOnce(notificationChan, task)
		err := t.handleTaskError(scope, task, notificationChan, err)
		if err != nil {
			task.attempt++
			if task.attempt >= t.config.TimerTaskMaxRetryCount() {
				scope.RecordTimer(metrics.TaskAttemptTimer, time.Duration(task.attempt))
				task.logger.Error("Critical error processing task, retrying.",
					tag.Error(err), tag.OperationCritical, tag.TaskType(task.task.GetTaskType()))
			}
		}
		return err
	}
	retryCondition := func(err error) bool {
		select {
		case <-t.shutdownCh:
			return false
		default:
			return true
		}
	}

	for {
		select {
		case <-t.shutdownCh:
			// this must return without ack
			return
		default:
			err = backoff.Retry(op, t.retryPolicy, retryCondition)
			if err == nil {
				t.ackTaskOnce(scope, task)
				return
			}
		}
	}
}

func (t *taskProcessor) processTaskOnce(
	notificationChan <-chan struct{},
	task *taskInfo,
) (metrics.Scope, error) {

	select {
	case <-notificationChan:
	default:
	}

	var scopeIdx int
	var err error

	startTime := t.timeSource.Now()
	scopeIdx, err = task.processor.process(task)

	//domainID := task.task.GetDomainID()

	newStartTime := t.timeSource.Now()

	var scope metrics.Scope
	var found bool
	if t.emitMetricsWithDomainTag {
		scope, found = t.domainMetricsScopeCache.Get(domainID, scopeIdx)

		if !found {
			domainTag, err := t.getDomainTagByID(domainID)
			scope = t.metricsClient.Scope(scopeIdx).Tagged(domainTag)
			// do not cache DomainUnknownTag
			if err == nil {
				t.domainMetricsScopeCache.Put(domainID, scopeIdx, scope)
			}
		}
	} else {
		scope = t.metricsClient.Scope(scopeIdx)
	}

	if task.shouldProcessTask {
		scope.RecordTimer(metrics.TestTaskProcessingLatency, time.Since(newStartTime))
		scope.IncCounter(metrics.TaskRequests)
		scope.RecordTimer(metrics.TaskProcessingLatency, time.Since(startTime))
	}

	return scope, err
}

func (t *taskProcessor) handleTaskError(
	scope metrics.Scope,
	taskInfo *taskInfo,
	notificationChan <-chan struct{},
	err error,
) error {

	if err == nil {
		return nil
	}

	if _, ok := err.(*workflow.EntityNotExistsError); ok {
		return nil
	}

	if transferTask, ok := taskInfo.task.(*persistence.TransferTaskInfo); ok &&
		transferTask.TaskType == persistence.TransferTaskTypeCloseExecution &&
		err == execution.ErrMissingWorkflowStartEvent &&
		t.config.EnableDropStuckTaskByDomainID(taskInfo.task.GetDomainID()) { // use domainID here to avoid accessing domainCache
		scope.IncCounter(metrics.TransferTaskMissingEventCounter)
		taskInfo.logger.Error("Drop close execution transfer task due to corrupted workflow history", tag.Error(err), tag.LifeCycleProcessingFailed)
		return nil
	}

	// this is a transient error
	if err == task.ErrTaskRetry {
		scope.IncCounter(metrics.TaskStandbyRetryCounter)
		select {
		case <-notificationChan:
		case <-t.shutdownCh:
		}
		return err
	}

	if err == task.ErrTaskDiscarded {
		scope.IncCounter(metrics.TaskDiscarded)
		err = nil
	}

	// this is a transient error
	// TODO remove this error check special case
	//  since the new task life cycle will not give up until task processed / verified
	if _, ok := err.(*workflow.DomainNotActiveError); ok {
		if t.timeSource.Now().Sub(taskInfo.startTime) > 2*cache.DomainCacheRefreshInterval {
			scope.IncCounter(metrics.TaskNotActiveCounter)
			return nil
		}

		return err
	}

	scope.IncCounter(metrics.TaskFailures)

	if _, ok := err.(*persistence.CurrentWorkflowConditionFailedError); ok {
		taskInfo.logger.Error("More than 2 workflow are running.", tag.Error(err), tag.LifeCycleProcessingFailed)
		return nil
	}

	taskInfo.logger.Error("Fail to process task", tag.Error(err), tag.LifeCycleProcessingFailed)
	return err
}

func (t *taskProcessor) ackTaskOnce(
	scope metrics.Scope,
	task *taskInfo,
) {

	task.processor.complete(task)
	if task.shouldProcessTask {
		scope.RecordTimer(metrics.TaskAttemptTimer, time.Duration(task.attempt))
		scope.RecordTimer(metrics.TaskLatency, time.Since(task.startTime))
		scope.RecordTimer(metrics.TaskQueueLatency, time.Since(task.task.GetVisibilityTimestamp()))
	}
}

func (t *taskProcessor) getDomainTagByID(domainID string) (metrics.Tag, error) {
	domainName, err := t.shard.GetDomainCache().GetDomainName(domainID)
	if err != nil {
		t.logger.Error("Unable to get domainName", tag.Error(err))
		return metrics.DomainUnknownTag(), err
	}
	return metrics.DomainTag(domainName), nil
}
