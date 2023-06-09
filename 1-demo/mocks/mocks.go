// Code generated by MockGen. DO NOT EDIT.
// Source: strings/strings.go

// Package mock_strings is a generated GoMock package.
package mock_strings

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStringer is a mock of Stringer interface.
type MockStringer struct {
	ctrl     *gomock.Controller
	recorder *MockStringerMockRecorder
}

// MockStringerMockRecorder is the mock recorder for MockStringer.
type MockStringerMockRecorder struct {
	mock *MockStringer
}

// NewMockStringer creates a new mock instance.
func NewMockStringer(ctrl *gomock.Controller) *MockStringer {
	mock := &MockStringer{ctrl: ctrl}
	mock.recorder = &MockStringerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStringer) EXPECT() *MockStringerMockRecorder {
	return m.recorder
}

// Reverse mocks base method.
func (m *MockStringer) Reverse() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reverse")
	ret0, _ := ret[0].(string)
	return ret0
}

// Reverse indicates an expected call of Reverse.
func (mr *MockStringerMockRecorder) Reverse() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reverse", reflect.TypeOf((*MockStringer)(nil).Reverse))
}
