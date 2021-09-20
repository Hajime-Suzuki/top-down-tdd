// Code generated by MockGen. DO NOT EDIT.
// Source: top-down-tdd/abstractions (interfaces: Players,Player)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	abstractions "top-down-tdd/abstractions"

	gomock "github.com/golang/mock/gomock"
)

// MockPlayers is a mock of Players interface.
type MockPlayers struct {
	ctrl     *gomock.Controller
	recorder *MockPlayersMockRecorder
}

// MockPlayersMockRecorder is the mock recorder for MockPlayers.
type MockPlayersMockRecorder struct {
	mock *MockPlayers
}

// NewMockPlayers creates a new mock instance.
func NewMockPlayers(ctrl *gomock.Controller) *MockPlayers {
	mock := &MockPlayers{ctrl: ctrl}
	mock.recorder = &MockPlayersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlayers) EXPECT() *MockPlayersMockRecorder {
	return m.recorder
}

// GetCurrentPlayer mocks base method.
func (m *MockPlayers) GetCurrentPlayer() abstractions.Player {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentPlayer")
	ret0, _ := ret[0].(abstractions.Player)
	return ret0
}

// GetCurrentPlayer indicates an expected call of GetCurrentPlayer.
func (mr *MockPlayersMockRecorder) GetCurrentPlayer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentPlayer", reflect.TypeOf((*MockPlayers)(nil).GetCurrentPlayer))
}

// GetPlayerByMark mocks base method.
func (m *MockPlayers) GetPlayerByMark(arg0 string) abstractions.Player {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlayerByMark", arg0)
	ret0, _ := ret[0].(abstractions.Player)
	return ret0
}

// GetPlayerByMark indicates an expected call of GetPlayerByMark.
func (mr *MockPlayersMockRecorder) GetPlayerByMark(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlayerByMark", reflect.TypeOf((*MockPlayers)(nil).GetPlayerByMark), arg0)
}

// GetPlayers mocks base method.
func (m *MockPlayers) GetPlayers() []abstractions.Player {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlayers")
	ret0, _ := ret[0].([]abstractions.Player)
	return ret0
}

// GetPlayers indicates an expected call of GetPlayers.
func (mr *MockPlayersMockRecorder) GetPlayers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlayers", reflect.TypeOf((*MockPlayers)(nil).GetPlayers))
}

// Next mocks base method.
func (m *MockPlayers) Next() abstractions.Players {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(abstractions.Players)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockPlayersMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockPlayers)(nil).Next))
}

// RegisterNewPlayer mocks base method.
func (m *MockPlayers) RegisterNewPlayer(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterNewPlayer", arg0)
}

// RegisterNewPlayer indicates an expected call of RegisterNewPlayer.
func (mr *MockPlayersMockRecorder) RegisterNewPlayer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterNewPlayer", reflect.TypeOf((*MockPlayers)(nil).RegisterNewPlayer), arg0)
}

// MockPlayer is a mock of Player interface.
type MockPlayer struct {
	ctrl     *gomock.Controller
	recorder *MockPlayerMockRecorder
}

// MockPlayerMockRecorder is the mock recorder for MockPlayer.
type MockPlayerMockRecorder struct {
	mock *MockPlayer
}

// NewMockPlayer creates a new mock instance.
func NewMockPlayer(ctrl *gomock.Controller) *MockPlayer {
	mock := &MockPlayer{ctrl: ctrl}
	mock.recorder = &MockPlayerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlayer) EXPECT() *MockPlayerMockRecorder {
	return m.recorder
}

// GetMark mocks base method.
func (m *MockPlayer) GetMark() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMark")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMark indicates an expected call of GetMark.
func (mr *MockPlayerMockRecorder) GetMark() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMark", reflect.TypeOf((*MockPlayer)(nil).GetMark))
}

// ShowName mocks base method.
func (m *MockPlayer) ShowName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ShowName indicates an expected call of ShowName.
func (mr *MockPlayerMockRecorder) ShowName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowName", reflect.TypeOf((*MockPlayer)(nil).ShowName))
}
