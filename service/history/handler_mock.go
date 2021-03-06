// The MIT License (MIT)
//
// Copyright (c) 2017-2020 Uber Technologies Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/cadence/service/history (interfaces: Handler)

// Package history is a generated GoMock package.
package history

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	health "github.com/uber/cadence/.gen/go/health"
	history "github.com/uber/cadence/.gen/go/history"
	replicator "github.com/uber/cadence/.gen/go/replicator"
	shared "github.com/uber/cadence/.gen/go/shared"
)

// MockHandler is a mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// CloseShard mocks base method
func (m *MockHandler) CloseShard(arg0 context.Context, arg1 *shared.CloseShardRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseShard", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseShard indicates an expected call of CloseShard
func (mr *MockHandlerMockRecorder) CloseShard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseShard", reflect.TypeOf((*MockHandler)(nil).CloseShard), arg0, arg1)
}

// DescribeHistoryHost mocks base method
func (m *MockHandler) DescribeHistoryHost(arg0 context.Context, arg1 *shared.DescribeHistoryHostRequest) (*shared.DescribeHistoryHostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeHistoryHost", arg0, arg1)
	ret0, _ := ret[0].(*shared.DescribeHistoryHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeHistoryHost indicates an expected call of DescribeHistoryHost
func (mr *MockHandlerMockRecorder) DescribeHistoryHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeHistoryHost", reflect.TypeOf((*MockHandler)(nil).DescribeHistoryHost), arg0, arg1)
}

// DescribeMutableState mocks base method
func (m *MockHandler) DescribeMutableState(arg0 context.Context, arg1 *history.DescribeMutableStateRequest) (*history.DescribeMutableStateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeMutableState", arg0, arg1)
	ret0, _ := ret[0].(*history.DescribeMutableStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMutableState indicates an expected call of DescribeMutableState
func (mr *MockHandlerMockRecorder) DescribeMutableState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMutableState", reflect.TypeOf((*MockHandler)(nil).DescribeMutableState), arg0, arg1)
}

// DescribeQueue mocks base method
func (m *MockHandler) DescribeQueue(arg0 context.Context, arg1 *shared.DescribeQueueRequest) (*shared.DescribeQueueResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeQueue", arg0, arg1)
	ret0, _ := ret[0].(*shared.DescribeQueueResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeQueue indicates an expected call of DescribeQueue
func (mr *MockHandlerMockRecorder) DescribeQueue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeQueue", reflect.TypeOf((*MockHandler)(nil).DescribeQueue), arg0, arg1)
}

// DescribeWorkflowExecution mocks base method
func (m *MockHandler) DescribeWorkflowExecution(arg0 context.Context, arg1 *history.DescribeWorkflowExecutionRequest) (*shared.DescribeWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(*shared.DescribeWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWorkflowExecution indicates an expected call of DescribeWorkflowExecution
func (mr *MockHandlerMockRecorder) DescribeWorkflowExecution(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).DescribeWorkflowExecution), arg0, arg1)
}

// GetDLQReplicationMessages mocks base method
func (m *MockHandler) GetDLQReplicationMessages(arg0 context.Context, arg1 *replicator.GetDLQReplicationMessagesRequest) (*replicator.GetDLQReplicationMessagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDLQReplicationMessages", arg0, arg1)
	ret0, _ := ret[0].(*replicator.GetDLQReplicationMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDLQReplicationMessages indicates an expected call of GetDLQReplicationMessages
func (mr *MockHandlerMockRecorder) GetDLQReplicationMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDLQReplicationMessages", reflect.TypeOf((*MockHandler)(nil).GetDLQReplicationMessages), arg0, arg1)
}

// GetMutableState mocks base method
func (m *MockHandler) GetMutableState(arg0 context.Context, arg1 *history.GetMutableStateRequest) (*history.GetMutableStateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMutableState", arg0, arg1)
	ret0, _ := ret[0].(*history.GetMutableStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMutableState indicates an expected call of GetMutableState
func (mr *MockHandlerMockRecorder) GetMutableState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMutableState", reflect.TypeOf((*MockHandler)(nil).GetMutableState), arg0, arg1)
}

// GetReplicationMessages mocks base method
func (m *MockHandler) GetReplicationMessages(arg0 context.Context, arg1 *replicator.GetReplicationMessagesRequest) (*replicator.GetReplicationMessagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicationMessages", arg0, arg1)
	ret0, _ := ret[0].(*replicator.GetReplicationMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReplicationMessages indicates an expected call of GetReplicationMessages
func (mr *MockHandlerMockRecorder) GetReplicationMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicationMessages", reflect.TypeOf((*MockHandler)(nil).GetReplicationMessages), arg0, arg1)
}

// Health mocks base method
func (m *MockHandler) Health(arg0 context.Context) (*health.HealthStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health", arg0)
	ret0, _ := ret[0].(*health.HealthStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Health indicates an expected call of Health
func (mr *MockHandlerMockRecorder) Health(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockHandler)(nil).Health), arg0)
}

// MergeDLQMessages mocks base method
func (m *MockHandler) MergeDLQMessages(arg0 context.Context, arg1 *replicator.MergeDLQMessagesRequest) (*replicator.MergeDLQMessagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MergeDLQMessages", arg0, arg1)
	ret0, _ := ret[0].(*replicator.MergeDLQMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MergeDLQMessages indicates an expected call of MergeDLQMessages
func (mr *MockHandlerMockRecorder) MergeDLQMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MergeDLQMessages", reflect.TypeOf((*MockHandler)(nil).MergeDLQMessages), arg0, arg1)
}

// NotifyFailoverMarkers mocks base method
func (m *MockHandler) NotifyFailoverMarkers(arg0 context.Context, arg1 *history.NotifyFailoverMarkersRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyFailoverMarkers", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyFailoverMarkers indicates an expected call of NotifyFailoverMarkers
func (mr *MockHandlerMockRecorder) NotifyFailoverMarkers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyFailoverMarkers", reflect.TypeOf((*MockHandler)(nil).NotifyFailoverMarkers), arg0, arg1)
}

// PollMutableState mocks base method
func (m *MockHandler) PollMutableState(arg0 context.Context, arg1 *history.PollMutableStateRequest) (*history.PollMutableStateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollMutableState", arg0, arg1)
	ret0, _ := ret[0].(*history.PollMutableStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollMutableState indicates an expected call of PollMutableState
func (mr *MockHandlerMockRecorder) PollMutableState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollMutableState", reflect.TypeOf((*MockHandler)(nil).PollMutableState), arg0, arg1)
}

// PurgeDLQMessages mocks base method
func (m *MockHandler) PurgeDLQMessages(arg0 context.Context, arg1 *replicator.PurgeDLQMessagesRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PurgeDLQMessages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PurgeDLQMessages indicates an expected call of PurgeDLQMessages
func (mr *MockHandlerMockRecorder) PurgeDLQMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeDLQMessages", reflect.TypeOf((*MockHandler)(nil).PurgeDLQMessages), arg0, arg1)
}

// QueryWorkflow mocks base method
func (m *MockHandler) QueryWorkflow(arg0 context.Context, arg1 *history.QueryWorkflowRequest) (*history.QueryWorkflowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryWorkflow", arg0, arg1)
	ret0, _ := ret[0].(*history.QueryWorkflowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryWorkflow indicates an expected call of QueryWorkflow
func (mr *MockHandlerMockRecorder) QueryWorkflow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryWorkflow", reflect.TypeOf((*MockHandler)(nil).QueryWorkflow), arg0, arg1)
}

// ReadDLQMessages mocks base method
func (m *MockHandler) ReadDLQMessages(arg0 context.Context, arg1 *replicator.ReadDLQMessagesRequest) (*replicator.ReadDLQMessagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDLQMessages", arg0, arg1)
	ret0, _ := ret[0].(*replicator.ReadDLQMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDLQMessages indicates an expected call of ReadDLQMessages
func (mr *MockHandlerMockRecorder) ReadDLQMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDLQMessages", reflect.TypeOf((*MockHandler)(nil).ReadDLQMessages), arg0, arg1)
}

// ReapplyEvents mocks base method
func (m *MockHandler) ReapplyEvents(arg0 context.Context, arg1 *history.ReapplyEventsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReapplyEvents", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReapplyEvents indicates an expected call of ReapplyEvents
func (mr *MockHandlerMockRecorder) ReapplyEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReapplyEvents", reflect.TypeOf((*MockHandler)(nil).ReapplyEvents), arg0, arg1)
}

// RecordActivityTaskHeartbeat mocks base method
func (m *MockHandler) RecordActivityTaskHeartbeat(arg0 context.Context, arg1 *history.RecordActivityTaskHeartbeatRequest) (*shared.RecordActivityTaskHeartbeatResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordActivityTaskHeartbeat", arg0, arg1)
	ret0, _ := ret[0].(*shared.RecordActivityTaskHeartbeatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordActivityTaskHeartbeat indicates an expected call of RecordActivityTaskHeartbeat
func (mr *MockHandlerMockRecorder) RecordActivityTaskHeartbeat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordActivityTaskHeartbeat", reflect.TypeOf((*MockHandler)(nil).RecordActivityTaskHeartbeat), arg0, arg1)
}

// RecordActivityTaskStarted mocks base method
func (m *MockHandler) RecordActivityTaskStarted(arg0 context.Context, arg1 *history.RecordActivityTaskStartedRequest) (*history.RecordActivityTaskStartedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordActivityTaskStarted", arg0, arg1)
	ret0, _ := ret[0].(*history.RecordActivityTaskStartedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordActivityTaskStarted indicates an expected call of RecordActivityTaskStarted
func (mr *MockHandlerMockRecorder) RecordActivityTaskStarted(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordActivityTaskStarted", reflect.TypeOf((*MockHandler)(nil).RecordActivityTaskStarted), arg0, arg1)
}

// RecordChildExecutionCompleted mocks base method
func (m *MockHandler) RecordChildExecutionCompleted(arg0 context.Context, arg1 *history.RecordChildExecutionCompletedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordChildExecutionCompleted", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordChildExecutionCompleted indicates an expected call of RecordChildExecutionCompleted
func (mr *MockHandlerMockRecorder) RecordChildExecutionCompleted(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordChildExecutionCompleted", reflect.TypeOf((*MockHandler)(nil).RecordChildExecutionCompleted), arg0, arg1)
}

// RecordDecisionTaskStarted mocks base method
func (m *MockHandler) RecordDecisionTaskStarted(arg0 context.Context, arg1 *history.RecordDecisionTaskStartedRequest) (*history.RecordDecisionTaskStartedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordDecisionTaskStarted", arg0, arg1)
	ret0, _ := ret[0].(*history.RecordDecisionTaskStartedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordDecisionTaskStarted indicates an expected call of RecordDecisionTaskStarted
func (mr *MockHandlerMockRecorder) RecordDecisionTaskStarted(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordDecisionTaskStarted", reflect.TypeOf((*MockHandler)(nil).RecordDecisionTaskStarted), arg0, arg1)
}

// RefreshWorkflowTasks mocks base method
func (m *MockHandler) RefreshWorkflowTasks(arg0 context.Context, arg1 *history.RefreshWorkflowTasksRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshWorkflowTasks", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshWorkflowTasks indicates an expected call of RefreshWorkflowTasks
func (mr *MockHandlerMockRecorder) RefreshWorkflowTasks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshWorkflowTasks", reflect.TypeOf((*MockHandler)(nil).RefreshWorkflowTasks), arg0, arg1)
}

// RemoveSignalMutableState mocks base method
func (m *MockHandler) RemoveSignalMutableState(arg0 context.Context, arg1 *history.RemoveSignalMutableStateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSignalMutableState", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSignalMutableState indicates an expected call of RemoveSignalMutableState
func (mr *MockHandlerMockRecorder) RemoveSignalMutableState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSignalMutableState", reflect.TypeOf((*MockHandler)(nil).RemoveSignalMutableState), arg0, arg1)
}

// RemoveTask mocks base method
func (m *MockHandler) RemoveTask(arg0 context.Context, arg1 *shared.RemoveTaskRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTask indicates an expected call of RemoveTask
func (mr *MockHandlerMockRecorder) RemoveTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTask", reflect.TypeOf((*MockHandler)(nil).RemoveTask), arg0, arg1)
}

// ReplicateEventsV2 mocks base method
func (m *MockHandler) ReplicateEventsV2(arg0 context.Context, arg1 *history.ReplicateEventsV2Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicateEventsV2", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplicateEventsV2 indicates an expected call of ReplicateEventsV2
func (mr *MockHandlerMockRecorder) ReplicateEventsV2(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicateEventsV2", reflect.TypeOf((*MockHandler)(nil).ReplicateEventsV2), arg0, arg1)
}

// RequestCancelWorkflowExecution mocks base method
func (m *MockHandler) RequestCancelWorkflowExecution(arg0 context.Context, arg1 *history.RequestCancelWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestCancelWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RequestCancelWorkflowExecution indicates an expected call of RequestCancelWorkflowExecution
func (mr *MockHandlerMockRecorder) RequestCancelWorkflowExecution(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestCancelWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).RequestCancelWorkflowExecution), arg0, arg1)
}

// ResetQueue mocks base method
func (m *MockHandler) ResetQueue(arg0 context.Context, arg1 *shared.ResetQueueRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetQueue", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetQueue indicates an expected call of ResetQueue
func (mr *MockHandlerMockRecorder) ResetQueue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetQueue", reflect.TypeOf((*MockHandler)(nil).ResetQueue), arg0, arg1)
}

// ResetStickyTaskList mocks base method
func (m *MockHandler) ResetStickyTaskList(arg0 context.Context, arg1 *history.ResetStickyTaskListRequest) (*history.ResetStickyTaskListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetStickyTaskList", arg0, arg1)
	ret0, _ := ret[0].(*history.ResetStickyTaskListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetStickyTaskList indicates an expected call of ResetStickyTaskList
func (mr *MockHandlerMockRecorder) ResetStickyTaskList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetStickyTaskList", reflect.TypeOf((*MockHandler)(nil).ResetStickyTaskList), arg0, arg1)
}

// ResetWorkflowExecution mocks base method
func (m *MockHandler) ResetWorkflowExecution(arg0 context.Context, arg1 *history.ResetWorkflowExecutionRequest) (*shared.ResetWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(*shared.ResetWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetWorkflowExecution indicates an expected call of ResetWorkflowExecution
func (mr *MockHandlerMockRecorder) ResetWorkflowExecution(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).ResetWorkflowExecution), arg0, arg1)
}

// RespondActivityTaskCanceled mocks base method
func (m *MockHandler) RespondActivityTaskCanceled(arg0 context.Context, arg1 *history.RespondActivityTaskCanceledRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondActivityTaskCanceled", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskCanceled indicates an expected call of RespondActivityTaskCanceled
func (mr *MockHandlerMockRecorder) RespondActivityTaskCanceled(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskCanceled", reflect.TypeOf((*MockHandler)(nil).RespondActivityTaskCanceled), arg0, arg1)
}

// RespondActivityTaskCompleted mocks base method
func (m *MockHandler) RespondActivityTaskCompleted(arg0 context.Context, arg1 *history.RespondActivityTaskCompletedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondActivityTaskCompleted", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskCompleted indicates an expected call of RespondActivityTaskCompleted
func (mr *MockHandlerMockRecorder) RespondActivityTaskCompleted(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskCompleted", reflect.TypeOf((*MockHandler)(nil).RespondActivityTaskCompleted), arg0, arg1)
}

// RespondActivityTaskFailed mocks base method
func (m *MockHandler) RespondActivityTaskFailed(arg0 context.Context, arg1 *history.RespondActivityTaskFailedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondActivityTaskFailed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskFailed indicates an expected call of RespondActivityTaskFailed
func (mr *MockHandlerMockRecorder) RespondActivityTaskFailed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskFailed", reflect.TypeOf((*MockHandler)(nil).RespondActivityTaskFailed), arg0, arg1)
}

// RespondDecisionTaskCompleted mocks base method
func (m *MockHandler) RespondDecisionTaskCompleted(arg0 context.Context, arg1 *history.RespondDecisionTaskCompletedRequest) (*history.RespondDecisionTaskCompletedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondDecisionTaskCompleted", arg0, arg1)
	ret0, _ := ret[0].(*history.RespondDecisionTaskCompletedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RespondDecisionTaskCompleted indicates an expected call of RespondDecisionTaskCompleted
func (mr *MockHandlerMockRecorder) RespondDecisionTaskCompleted(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondDecisionTaskCompleted", reflect.TypeOf((*MockHandler)(nil).RespondDecisionTaskCompleted), arg0, arg1)
}

// RespondDecisionTaskFailed mocks base method
func (m *MockHandler) RespondDecisionTaskFailed(arg0 context.Context, arg1 *history.RespondDecisionTaskFailedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondDecisionTaskFailed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondDecisionTaskFailed indicates an expected call of RespondDecisionTaskFailed
func (mr *MockHandlerMockRecorder) RespondDecisionTaskFailed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondDecisionTaskFailed", reflect.TypeOf((*MockHandler)(nil).RespondDecisionTaskFailed), arg0, arg1)
}

// ScheduleDecisionTask mocks base method
func (m *MockHandler) ScheduleDecisionTask(arg0 context.Context, arg1 *history.ScheduleDecisionTaskRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleDecisionTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScheduleDecisionTask indicates an expected call of ScheduleDecisionTask
func (mr *MockHandlerMockRecorder) ScheduleDecisionTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleDecisionTask", reflect.TypeOf((*MockHandler)(nil).ScheduleDecisionTask), arg0, arg1)
}

// SignalWithStartWorkflowExecution mocks base method
func (m *MockHandler) SignalWithStartWorkflowExecution(arg0 context.Context, arg1 *history.SignalWithStartWorkflowExecutionRequest) (*shared.StartWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignalWithStartWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(*shared.StartWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignalWithStartWorkflowExecution indicates an expected call of SignalWithStartWorkflowExecution
func (mr *MockHandlerMockRecorder) SignalWithStartWorkflowExecution(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignalWithStartWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).SignalWithStartWorkflowExecution), arg0, arg1)
}

// SignalWorkflowExecution mocks base method
func (m *MockHandler) SignalWorkflowExecution(arg0 context.Context, arg1 *history.SignalWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignalWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignalWorkflowExecution indicates an expected call of SignalWorkflowExecution
func (mr *MockHandlerMockRecorder) SignalWorkflowExecution(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignalWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).SignalWorkflowExecution), arg0, arg1)
}

// StartWorkflowExecution mocks base method
func (m *MockHandler) StartWorkflowExecution(arg0 context.Context, arg1 *history.StartWorkflowExecutionRequest) (*shared.StartWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(*shared.StartWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartWorkflowExecution indicates an expected call of StartWorkflowExecution
func (mr *MockHandlerMockRecorder) StartWorkflowExecution(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).StartWorkflowExecution), arg0, arg1)
}

// SyncActivity mocks base method
func (m *MockHandler) SyncActivity(arg0 context.Context, arg1 *history.SyncActivityRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncActivity", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncActivity indicates an expected call of SyncActivity
func (mr *MockHandlerMockRecorder) SyncActivity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncActivity", reflect.TypeOf((*MockHandler)(nil).SyncActivity), arg0, arg1)
}

// SyncShardStatus mocks base method
func (m *MockHandler) SyncShardStatus(arg0 context.Context, arg1 *history.SyncShardStatusRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncShardStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncShardStatus indicates an expected call of SyncShardStatus
func (mr *MockHandlerMockRecorder) SyncShardStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncShardStatus", reflect.TypeOf((*MockHandler)(nil).SyncShardStatus), arg0, arg1)
}

// TerminateWorkflowExecution mocks base method
func (m *MockHandler) TerminateWorkflowExecution(arg0 context.Context, arg1 *history.TerminateWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TerminateWorkflowExecution indicates an expected call of TerminateWorkflowExecution
func (mr *MockHandlerMockRecorder) TerminateWorkflowExecution(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateWorkflowExecution", reflect.TypeOf((*MockHandler)(nil).TerminateWorkflowExecution), arg0, arg1)
}
