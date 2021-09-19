// Code generated by MockGen. DO NOT EDIT.
// Source: top-down-tdd/abstractions (interfaces: Presenter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	abstractions "top-down-tdd/abstractions"

	gomock "github.com/golang/mock/gomock"
)

// MockPresenter is a mock of Presenter interface.
type MockPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockPresenterMockRecorder
}

// MockPresenterMockRecorder is the mock recorder for MockPresenter.
type MockPresenterMockRecorder struct {
	mock *MockPresenter
}

// NewMockPresenter creates a new mock instance.
func NewMockPresenter(ctrl *gomock.Controller) *MockPresenter {
	mock := &MockPresenter{ctrl: ctrl}
	mock.recorder = &MockPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPresenter) EXPECT() *MockPresenterMockRecorder {
	return m.recorder
}

// Dispay mocks base method.
func (m *MockPresenter) Dispay(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Dispay", arg0)
}

// Dispay indicates an expected call of Dispay.
func (mr *MockPresenterMockRecorder) Dispay(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispay", reflect.TypeOf((*MockPresenter)(nil).Dispay), arg0)
}

// ShowMessage mocks base method.
func (m *MockPresenter) ShowMessage(arg0 abstractions.Player) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowMessage", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// ShowMessage indicates an expected call of ShowMessage.
func (mr *MockPresenterMockRecorder) ShowMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowMessage", reflect.TypeOf((*MockPresenter)(nil).ShowMessage), arg0)
}
