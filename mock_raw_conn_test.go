// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apernet/quic-go (interfaces: RawConn)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/apernet/quic-go -destination mock_raw_conn_test.go github.com/apernet/quic-go RawConn
//
// Package quic is a generated GoMock package.
package quic

import (
	net "net"
	reflect "reflect"
	time "time"

	protocol "github.com/apernet/quic-go/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockRawConn is a mock of RawConn interface.
type MockRawConn struct {
	ctrl     *gomock.Controller
	recorder *MockRawConnMockRecorder
}

// MockRawConnMockRecorder is the mock recorder for MockRawConn.
type MockRawConnMockRecorder struct {
	mock *MockRawConn
}

// NewMockRawConn creates a new mock instance.
func NewMockRawConn(ctrl *gomock.Controller) *MockRawConn {
	mock := &MockRawConn{ctrl: ctrl}
	mock.recorder = &MockRawConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRawConn) EXPECT() *MockRawConnMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRawConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRawConnMockRecorder) Close() *RawConnCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRawConn)(nil).Close))
	return &RawConnCloseCall{Call: call}
}

// RawConnCloseCall wrap *gomock.Call
type RawConnCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RawConnCloseCall) Return(arg0 error) *RawConnCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RawConnCloseCall) Do(f func() error) *RawConnCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RawConnCloseCall) DoAndReturn(f func() error) *RawConnCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LocalAddr mocks base method.
func (m *MockRawConn) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockRawConnMockRecorder) LocalAddr() *RawConnLocalAddrCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockRawConn)(nil).LocalAddr))
	return &RawConnLocalAddrCall{Call: call}
}

// RawConnLocalAddrCall wrap *gomock.Call
type RawConnLocalAddrCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RawConnLocalAddrCall) Return(arg0 net.Addr) *RawConnLocalAddrCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RawConnLocalAddrCall) Do(f func() net.Addr) *RawConnLocalAddrCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RawConnLocalAddrCall) DoAndReturn(f func() net.Addr) *RawConnLocalAddrCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadPacket mocks base method.
func (m *MockRawConn) ReadPacket() (receivedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPacket")
	ret0, _ := ret[0].(receivedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPacket indicates an expected call of ReadPacket.
func (mr *MockRawConnMockRecorder) ReadPacket() *RawConnReadPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPacket", reflect.TypeOf((*MockRawConn)(nil).ReadPacket))
	return &RawConnReadPacketCall{Call: call}
}

// RawConnReadPacketCall wrap *gomock.Call
type RawConnReadPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RawConnReadPacketCall) Return(arg0 receivedPacket, arg1 error) *RawConnReadPacketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RawConnReadPacketCall) Do(f func() (receivedPacket, error)) *RawConnReadPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RawConnReadPacketCall) DoAndReturn(f func() (receivedPacket, error)) *RawConnReadPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetReadDeadline mocks base method.
func (m *MockRawConn) SetReadDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockRawConnMockRecorder) SetReadDeadline(arg0 any) *RawConnSetReadDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockRawConn)(nil).SetReadDeadline), arg0)
	return &RawConnSetReadDeadlineCall{Call: call}
}

// RawConnSetReadDeadlineCall wrap *gomock.Call
type RawConnSetReadDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RawConnSetReadDeadlineCall) Return(arg0 error) *RawConnSetReadDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RawConnSetReadDeadlineCall) Do(f func(time.Time) error) *RawConnSetReadDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RawConnSetReadDeadlineCall) DoAndReturn(f func(time.Time) error) *RawConnSetReadDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WritePacket mocks base method.
func (m *MockRawConn) WritePacket(arg0 []byte, arg1 net.Addr, arg2 []byte, arg3 uint16, arg4 protocol.ECN) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WritePacket", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WritePacket indicates an expected call of WritePacket.
func (mr *MockRawConnMockRecorder) WritePacket(arg0, arg1, arg2, arg3, arg4 any) *RawConnWritePacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePacket", reflect.TypeOf((*MockRawConn)(nil).WritePacket), arg0, arg1, arg2, arg3, arg4)
	return &RawConnWritePacketCall{Call: call}
}

// RawConnWritePacketCall wrap *gomock.Call
type RawConnWritePacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RawConnWritePacketCall) Return(arg0 int, arg1 error) *RawConnWritePacketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RawConnWritePacketCall) Do(f func([]byte, net.Addr, []byte, uint16, protocol.ECN) (int, error)) *RawConnWritePacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RawConnWritePacketCall) DoAndReturn(f func([]byte, net.Addr, []byte, uint16, protocol.ECN) (int, error)) *RawConnWritePacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// capabilities mocks base method.
func (m *MockRawConn) capabilities() connCapabilities {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "capabilities")
	ret0, _ := ret[0].(connCapabilities)
	return ret0
}

// capabilities indicates an expected call of capabilities.
func (mr *MockRawConnMockRecorder) capabilities() *RawConncapabilitiesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "capabilities", reflect.TypeOf((*MockRawConn)(nil).capabilities))
	return &RawConncapabilitiesCall{Call: call}
}

// RawConncapabilitiesCall wrap *gomock.Call
type RawConncapabilitiesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RawConncapabilitiesCall) Return(arg0 connCapabilities) *RawConncapabilitiesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RawConncapabilitiesCall) Do(f func() connCapabilities) *RawConncapabilitiesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RawConncapabilitiesCall) DoAndReturn(f func() connCapabilities) *RawConncapabilitiesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
