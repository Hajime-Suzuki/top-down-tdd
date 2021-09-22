// Code generated by MockGen. DO NOT EDIT.
// Source: top-down-tdd/abstractions (interfaces: InputHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInputHandler is a mock of InputHandler interface.
type MockInputHandler struct {
	ctrl     *gomock.Controller
	recorder *MockInputHandlerMockRecorder
}

// MockInputHandlerMockRecorder is the mock recorder for MockInputHandler.
type MockInputHandlerMockRecorder struct {
	mock *MockInputHandler
}

// NewMockInputHandler creates a new mock instance.
func NewMockInputHandler(ctrl *gomock.Controller) *MockInputHandler {
	mock := &MockInputHandler{ctrl: ctrl}
	mock.recorder = &MockInputHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInputHandler) EXPECT() *MockInputHandlerMockRecorder {
	return m.recorder
}

// GetUserInput mocks base method.
func (m *MockInputHandler) GetUserInput() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInput")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUserInput indicates an expected call of GetUserInput.
func (mr *MockInputHandlerMockRecorder) GetUserInput() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInput", reflect.TypeOf((*MockInputHandler)(nil).GetUserInput))
}
