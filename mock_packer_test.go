// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apernet/quic-go (interfaces: Packer)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/apernet/quic-go -destination mock_packer_test.go github.com/apernet/quic-go Packer
//
// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	ackhandler "github.com/apernet/quic-go/internal/ackhandler"
	protocol "github.com/apernet/quic-go/internal/protocol"
	qerr "github.com/apernet/quic-go/internal/qerr"
	gomock "go.uber.org/mock/gomock"
)

// MockPacker is a mock of Packer interface.
type MockPacker struct {
	ctrl     *gomock.Controller
	recorder *MockPackerMockRecorder
}

// MockPackerMockRecorder is the mock recorder for MockPacker.
type MockPackerMockRecorder struct {
	mock *MockPacker
}

// NewMockPacker creates a new mock instance.
func NewMockPacker(ctrl *gomock.Controller) *MockPacker {
	mock := &MockPacker{ctrl: ctrl}
	mock.recorder = &MockPackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacker) EXPECT() *MockPackerMockRecorder {
	return m.recorder
}

// AppendPacket mocks base method.
func (m *MockPacker) AppendPacket(arg0 *packetBuffer, arg1 protocol.ByteCount, arg2 protocol.VersionNumber) (shortHeaderPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendPacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(shortHeaderPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppendPacket indicates an expected call of AppendPacket.
func (mr *MockPackerMockRecorder) AppendPacket(arg0, arg1, arg2 any) *PackerAppendPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendPacket", reflect.TypeOf((*MockPacker)(nil).AppendPacket), arg0, arg1, arg2)
	return &PackerAppendPacketCall{Call: call}
}

// PackerAppendPacketCall wrap *gomock.Call
type PackerAppendPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PackerAppendPacketCall) Return(arg0 shortHeaderPacket, arg1 error) *PackerAppendPacketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PackerAppendPacketCall) Do(f func(*packetBuffer, protocol.ByteCount, protocol.VersionNumber) (shortHeaderPacket, error)) *PackerAppendPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PackerAppendPacketCall) DoAndReturn(f func(*packetBuffer, protocol.ByteCount, protocol.VersionNumber) (shortHeaderPacket, error)) *PackerAppendPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MaybePackProbePacket mocks base method.
func (m *MockPacker) MaybePackProbePacket(arg0 protocol.EncryptionLevel, arg1 protocol.ByteCount, arg2 protocol.VersionNumber) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaybePackProbePacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaybePackProbePacket indicates an expected call of MaybePackProbePacket.
func (mr *MockPackerMockRecorder) MaybePackProbePacket(arg0, arg1, arg2 any) *PackerMaybePackProbePacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybePackProbePacket", reflect.TypeOf((*MockPacker)(nil).MaybePackProbePacket), arg0, arg1, arg2)
	return &PackerMaybePackProbePacketCall{Call: call}
}

// PackerMaybePackProbePacketCall wrap *gomock.Call
type PackerMaybePackProbePacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PackerMaybePackProbePacketCall) Return(arg0 *coalescedPacket, arg1 error) *PackerMaybePackProbePacketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PackerMaybePackProbePacketCall) Do(f func(protocol.EncryptionLevel, protocol.ByteCount, protocol.VersionNumber) (*coalescedPacket, error)) *PackerMaybePackProbePacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PackerMaybePackProbePacketCall) DoAndReturn(f func(protocol.EncryptionLevel, protocol.ByteCount, protocol.VersionNumber) (*coalescedPacket, error)) *PackerMaybePackProbePacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackAckOnlyPacket mocks base method.
func (m *MockPacker) PackAckOnlyPacket(arg0 protocol.ByteCount, arg1 protocol.VersionNumber) (shortHeaderPacket, *packetBuffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackAckOnlyPacket", arg0, arg1)
	ret0, _ := ret[0].(shortHeaderPacket)
	ret1, _ := ret[1].(*packetBuffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PackAckOnlyPacket indicates an expected call of PackAckOnlyPacket.
func (mr *MockPackerMockRecorder) PackAckOnlyPacket(arg0, arg1 any) *PackerPackAckOnlyPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackAckOnlyPacket", reflect.TypeOf((*MockPacker)(nil).PackAckOnlyPacket), arg0, arg1)
	return &PackerPackAckOnlyPacketCall{Call: call}
}

// PackerPackAckOnlyPacketCall wrap *gomock.Call
type PackerPackAckOnlyPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PackerPackAckOnlyPacketCall) Return(arg0 shortHeaderPacket, arg1 *packetBuffer, arg2 error) *PackerPackAckOnlyPacketCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PackerPackAckOnlyPacketCall) Do(f func(protocol.ByteCount, protocol.VersionNumber) (shortHeaderPacket, *packetBuffer, error)) *PackerPackAckOnlyPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PackerPackAckOnlyPacketCall) DoAndReturn(f func(protocol.ByteCount, protocol.VersionNumber) (shortHeaderPacket, *packetBuffer, error)) *PackerPackAckOnlyPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackApplicationClose mocks base method.
func (m *MockPacker) PackApplicationClose(arg0 *qerr.ApplicationError, arg1 protocol.ByteCount, arg2 protocol.VersionNumber) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackApplicationClose", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackApplicationClose indicates an expected call of PackApplicationClose.
func (mr *MockPackerMockRecorder) PackApplicationClose(arg0, arg1, arg2 any) *PackerPackApplicationCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackApplicationClose", reflect.TypeOf((*MockPacker)(nil).PackApplicationClose), arg0, arg1, arg2)
	return &PackerPackApplicationCloseCall{Call: call}
}

// PackerPackApplicationCloseCall wrap *gomock.Call
type PackerPackApplicationCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PackerPackApplicationCloseCall) Return(arg0 *coalescedPacket, arg1 error) *PackerPackApplicationCloseCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PackerPackApplicationCloseCall) Do(f func(*qerr.ApplicationError, protocol.ByteCount, protocol.VersionNumber) (*coalescedPacket, error)) *PackerPackApplicationCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PackerPackApplicationCloseCall) DoAndReturn(f func(*qerr.ApplicationError, protocol.ByteCount, protocol.VersionNumber) (*coalescedPacket, error)) *PackerPackApplicationCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackCoalescedPacket mocks base method.
func (m *MockPacker) PackCoalescedPacket(arg0 bool, arg1 protocol.ByteCount, arg2 protocol.VersionNumber) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackCoalescedPacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackCoalescedPacket indicates an expected call of PackCoalescedPacket.
func (mr *MockPackerMockRecorder) PackCoalescedPacket(arg0, arg1, arg2 any) *PackerPackCoalescedPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackCoalescedPacket", reflect.TypeOf((*MockPacker)(nil).PackCoalescedPacket), arg0, arg1, arg2)
	return &PackerPackCoalescedPacketCall{Call: call}
}

// PackerPackCoalescedPacketCall wrap *gomock.Call
type PackerPackCoalescedPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PackerPackCoalescedPacketCall) Return(arg0 *coalescedPacket, arg1 error) *PackerPackCoalescedPacketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PackerPackCoalescedPacketCall) Do(f func(bool, protocol.ByteCount, protocol.VersionNumber) (*coalescedPacket, error)) *PackerPackCoalescedPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PackerPackCoalescedPacketCall) DoAndReturn(f func(bool, protocol.ByteCount, protocol.VersionNumber) (*coalescedPacket, error)) *PackerPackCoalescedPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackConnectionClose mocks base method.
func (m *MockPacker) PackConnectionClose(arg0 *qerr.TransportError, arg1 protocol.ByteCount, arg2 protocol.VersionNumber) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackConnectionClose", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackConnectionClose indicates an expected call of PackConnectionClose.
func (mr *MockPackerMockRecorder) PackConnectionClose(arg0, arg1, arg2 any) *PackerPackConnectionCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackConnectionClose", reflect.TypeOf((*MockPacker)(nil).PackConnectionClose), arg0, arg1, arg2)
	return &PackerPackConnectionCloseCall{Call: call}
}

// PackerPackConnectionCloseCall wrap *gomock.Call
type PackerPackConnectionCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PackerPackConnectionCloseCall) Return(arg0 *coalescedPacket, arg1 error) *PackerPackConnectionCloseCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PackerPackConnectionCloseCall) Do(f func(*qerr.TransportError, protocol.ByteCount, protocol.VersionNumber) (*coalescedPacket, error)) *PackerPackConnectionCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PackerPackConnectionCloseCall) DoAndReturn(f func(*qerr.TransportError, protocol.ByteCount, protocol.VersionNumber) (*coalescedPacket, error)) *PackerPackConnectionCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackMTUProbePacket mocks base method.
func (m *MockPacker) PackMTUProbePacket(arg0 ackhandler.Frame, arg1 protocol.ByteCount, arg2 protocol.VersionNumber) (shortHeaderPacket, *packetBuffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackMTUProbePacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(shortHeaderPacket)
	ret1, _ := ret[1].(*packetBuffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PackMTUProbePacket indicates an expected call of PackMTUProbePacket.
func (mr *MockPackerMockRecorder) PackMTUProbePacket(arg0, arg1, arg2 any) *PackerPackMTUProbePacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackMTUProbePacket", reflect.TypeOf((*MockPacker)(nil).PackMTUProbePacket), arg0, arg1, arg2)
	return &PackerPackMTUProbePacketCall{Call: call}
}

// PackerPackMTUProbePacketCall wrap *gomock.Call
type PackerPackMTUProbePacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PackerPackMTUProbePacketCall) Return(arg0 shortHeaderPacket, arg1 *packetBuffer, arg2 error) *PackerPackMTUProbePacketCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PackerPackMTUProbePacketCall) Do(f func(ackhandler.Frame, protocol.ByteCount, protocol.VersionNumber) (shortHeaderPacket, *packetBuffer, error)) *PackerPackMTUProbePacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PackerPackMTUProbePacketCall) DoAndReturn(f func(ackhandler.Frame, protocol.ByteCount, protocol.VersionNumber) (shortHeaderPacket, *packetBuffer, error)) *PackerPackMTUProbePacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetToken mocks base method.
func (m *MockPacker) SetToken(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetToken", arg0)
}

// SetToken indicates an expected call of SetToken.
func (mr *MockPackerMockRecorder) SetToken(arg0 any) *PackerSetTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetToken", reflect.TypeOf((*MockPacker)(nil).SetToken), arg0)
	return &PackerSetTokenCall{Call: call}
}

// PackerSetTokenCall wrap *gomock.Call
type PackerSetTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PackerSetTokenCall) Return() *PackerSetTokenCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PackerSetTokenCall) Do(f func([]byte)) *PackerSetTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PackerSetTokenCall) DoAndReturn(f func([]byte)) *PackerSetTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
