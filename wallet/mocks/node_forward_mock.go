// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/go-wallet/wallet (interfaces: NodeForward)

// Package mocks is a generated GoMock package.
package mocks

import (
	wallet "code.vegaprotocol.io/go-wallet/wallet"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNodeForward is a mock of NodeForward interface
type MockNodeForward struct {
	ctrl     *gomock.Controller
	recorder *MockNodeForwardMockRecorder
}

// MockNodeForwardMockRecorder is the mock recorder for MockNodeForward
type MockNodeForwardMockRecorder struct {
	mock *MockNodeForward
}

// NewMockNodeForward creates a new mock instance
func NewMockNodeForward(ctrl *gomock.Controller) *MockNodeForward {
	mock := &MockNodeForward{ctrl: ctrl}
	mock.recorder = &MockNodeForwardMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeForward) EXPECT() *MockNodeForwardMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockNodeForward) Send(arg0 context.Context, arg1 *wallet.SignedBundle) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockNodeForwardMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNodeForward)(nil).Send), arg0, arg1)
}