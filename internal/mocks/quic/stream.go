// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apernet/quic-go (interfaces: Stream)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package mockquic -destination quic/stream.go github.com/apernet/quic-go Stream
//
// Package mockquic is a generated GoMock package.
package mockquic

import (
	context "context"
	reflect "reflect"
	time "time"

	protocol "github.com/apernet/quic-go/internal/protocol"
	qerr "github.com/apernet/quic-go/internal/qerr"
	gomock "go.uber.org/mock/gomock"
)

// MockStream is a mock of Stream interface.
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *MockStreamMockRecorder
}

// MockStreamMockRecorder is the mock recorder for MockStream.
type MockStreamMockRecorder struct {
	mock *MockStream
}

// NewMockStream creates a new mock instance.
func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &MockStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStream) EXPECT() *MockStreamMockRecorder {
	return m.recorder
}

// CancelRead mocks base method.
func (m *MockStream) CancelRead(arg0 qerr.StreamErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelRead", arg0)
}

// CancelRead indicates an expected call of CancelRead.
func (mr *MockStreamMockRecorder) CancelRead(arg0 any) *StreamCancelReadCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRead", reflect.TypeOf((*MockStream)(nil).CancelRead), arg0)
	return &StreamCancelReadCall{Call: call}
}

// StreamCancelReadCall wrap *gomock.Call
type StreamCancelReadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamCancelReadCall) Return() *StreamCancelReadCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamCancelReadCall) Do(f func(qerr.StreamErrorCode)) *StreamCancelReadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamCancelReadCall) DoAndReturn(f func(qerr.StreamErrorCode)) *StreamCancelReadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CancelWrite mocks base method.
func (m *MockStream) CancelWrite(arg0 qerr.StreamErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelWrite", arg0)
}

// CancelWrite indicates an expected call of CancelWrite.
func (mr *MockStreamMockRecorder) CancelWrite(arg0 any) *StreamCancelWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelWrite", reflect.TypeOf((*MockStream)(nil).CancelWrite), arg0)
	return &StreamCancelWriteCall{Call: call}
}

// StreamCancelWriteCall wrap *gomock.Call
type StreamCancelWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamCancelWriteCall) Return() *StreamCancelWriteCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamCancelWriteCall) Do(f func(qerr.StreamErrorCode)) *StreamCancelWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamCancelWriteCall) DoAndReturn(f func(qerr.StreamErrorCode)) *StreamCancelWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockStream) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStreamMockRecorder) Close() *StreamCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStream)(nil).Close))
	return &StreamCloseCall{Call: call}
}

// StreamCloseCall wrap *gomock.Call
type StreamCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamCloseCall) Return(arg0 error) *StreamCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamCloseCall) Do(f func() error) *StreamCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamCloseCall) DoAndReturn(f func() error) *StreamCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Context mocks base method.
func (m *MockStream) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockStreamMockRecorder) Context() *StreamContextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockStream)(nil).Context))
	return &StreamContextCall{Call: call}
}

// StreamContextCall wrap *gomock.Call
type StreamContextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamContextCall) Return(arg0 context.Context) *StreamContextCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamContextCall) Do(f func() context.Context) *StreamContextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamContextCall) DoAndReturn(f func() context.Context) *StreamContextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Read mocks base method.
func (m *MockStream) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockStreamMockRecorder) Read(arg0 any) *StreamReadCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockStream)(nil).Read), arg0)
	return &StreamReadCall{Call: call}
}

// StreamReadCall wrap *gomock.Call
type StreamReadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamReadCall) Return(arg0 int, arg1 error) *StreamReadCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamReadCall) Do(f func([]byte) (int, error)) *StreamReadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamReadCall) DoAndReturn(f func([]byte) (int, error)) *StreamReadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetDeadline mocks base method.
func (m *MockStream) SetDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeadline indicates an expected call of SetDeadline.
func (mr *MockStreamMockRecorder) SetDeadline(arg0 any) *StreamSetDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadline", reflect.TypeOf((*MockStream)(nil).SetDeadline), arg0)
	return &StreamSetDeadlineCall{Call: call}
}

// StreamSetDeadlineCall wrap *gomock.Call
type StreamSetDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamSetDeadlineCall) Return(arg0 error) *StreamSetDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamSetDeadlineCall) Do(f func(time.Time) error) *StreamSetDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamSetDeadlineCall) DoAndReturn(f func(time.Time) error) *StreamSetDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetReadDeadline mocks base method.
func (m *MockStream) SetReadDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockStreamMockRecorder) SetReadDeadline(arg0 any) *StreamSetReadDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockStream)(nil).SetReadDeadline), arg0)
	return &StreamSetReadDeadlineCall{Call: call}
}

// StreamSetReadDeadlineCall wrap *gomock.Call
type StreamSetReadDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamSetReadDeadlineCall) Return(arg0 error) *StreamSetReadDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamSetReadDeadlineCall) Do(f func(time.Time) error) *StreamSetReadDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamSetReadDeadlineCall) DoAndReturn(f func(time.Time) error) *StreamSetReadDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetWriteDeadline mocks base method.
func (m *MockStream) SetWriteDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline.
func (mr *MockStreamMockRecorder) SetWriteDeadline(arg0 any) *StreamSetWriteDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockStream)(nil).SetWriteDeadline), arg0)
	return &StreamSetWriteDeadlineCall{Call: call}
}

// StreamSetWriteDeadlineCall wrap *gomock.Call
type StreamSetWriteDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamSetWriteDeadlineCall) Return(arg0 error) *StreamSetWriteDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamSetWriteDeadlineCall) Do(f func(time.Time) error) *StreamSetWriteDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamSetWriteDeadlineCall) DoAndReturn(f func(time.Time) error) *StreamSetWriteDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StreamID mocks base method.
func (m *MockStream) StreamID() protocol.StreamID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamID")
	ret0, _ := ret[0].(protocol.StreamID)
	return ret0
}

// StreamID indicates an expected call of StreamID.
func (mr *MockStreamMockRecorder) StreamID() *StreamStreamIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamID", reflect.TypeOf((*MockStream)(nil).StreamID))
	return &StreamStreamIDCall{Call: call}
}

// StreamStreamIDCall wrap *gomock.Call
type StreamStreamIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamStreamIDCall) Return(arg0 protocol.StreamID) *StreamStreamIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamStreamIDCall) Do(f func() protocol.StreamID) *StreamStreamIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamStreamIDCall) DoAndReturn(f func() protocol.StreamID) *StreamStreamIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Write mocks base method.
func (m *MockStream) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockStreamMockRecorder) Write(arg0 any) *StreamWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockStream)(nil).Write), arg0)
	return &StreamWriteCall{Call: call}
}

// StreamWriteCall wrap *gomock.Call
type StreamWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamWriteCall) Return(arg0 int, arg1 error) *StreamWriteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamWriteCall) Do(f func([]byte) (int, error)) *StreamWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamWriteCall) DoAndReturn(f func([]byte) (int, error)) *StreamWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
