// Code generated by MockGen. DO NOT EDIT.
// Source: top-down-tdd/abstractions (interfaces: Board)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	abstractions "top-down-tdd/abstractions"

	gomock "github.com/golang/mock/gomock"
)

// MockBoard is a mock of Board interface.
type MockBoard struct {
	ctrl     *gomock.Controller
	recorder *MockBoardMockRecorder
}

// MockBoardMockRecorder is the mock recorder for MockBoard.
type MockBoardMockRecorder struct {
	mock *MockBoard
}

// NewMockBoard creates a new mock instance.
func NewMockBoard(ctrl *gomock.Controller) *MockBoard {
	mock := &MockBoard{ctrl: ctrl}
	mock.recorder = &MockBoardMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBoard) EXPECT() *MockBoardMockRecorder {
	return m.recorder
}

// GetWinner mocks base method.
func (m *MockBoard) GetWinner() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWinner")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWinner indicates an expected call of GetWinner.
func (mr *MockBoardMockRecorder) GetWinner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWinner", reflect.TypeOf((*MockBoard)(nil).GetWinner))
}

// IsOver mocks base method.
func (m *MockBoard) IsOver() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOver")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOver indicates an expected call of IsOver.
func (mr *MockBoardMockRecorder) IsOver() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOver", reflect.TypeOf((*MockBoard)(nil).IsOver))
}

// Show mocks base method.
func (m *MockBoard) Show() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show")
	ret0, _ := ret[0].(string)
	return ret0
}

// Show indicates an expected call of Show.
func (mr *MockBoardMockRecorder) Show() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockBoard)(nil).Show))
}

// Update mocks base method.
func (m *MockBoard) Update(arg0, arg1 string) (abstractions.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(abstractions.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBoardMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBoard)(nil).Update), arg0, arg1)
}
