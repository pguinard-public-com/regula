// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fugue/regula/v3/pkg/loader (interfaces: InputPath)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	loader "github.com/fugue/regula/v3/pkg/loader"
	gomock "github.com/golang/mock/gomock"
)

// MockInputPath is a mock of InputPath interface.
type MockInputPath struct {
	ctrl     *gomock.Controller
	recorder *MockInputPathMockRecorder
}

// MockInputPathMockRecorder is the mock recorder for MockInputPath.
type MockInputPathMockRecorder struct {
	mock *MockInputPath
}

// NewMockInputPath creates a new mock instance.
func NewMockInputPath(ctrl *gomock.Controller) *MockInputPath {
	mock := &MockInputPath{ctrl: ctrl}
	mock.recorder = &MockInputPathMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInputPath) EXPECT() *MockInputPathMockRecorder {
	return m.recorder
}

// DetectType mocks base method.
func (m *MockInputPath) DetectType(arg0 loader.ConfigurationDetector, arg1 loader.DetectOptions) (loader.IACConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectType", arg0, arg1)
	ret0, _ := ret[0].(loader.IACConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectType indicates an expected call of DetectType.
func (mr *MockInputPathMockRecorder) DetectType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectType", reflect.TypeOf((*MockInputPath)(nil).DetectType), arg0, arg1)
}

// IsDir mocks base method.
func (m *MockInputPath) IsDir() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDir")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDir indicates an expected call of IsDir.
func (mr *MockInputPathMockRecorder) IsDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDir", reflect.TypeOf((*MockInputPath)(nil).IsDir))
}

// Name mocks base method.
func (m *MockInputPath) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockInputPathMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockInputPath)(nil).Name))
}

// Path mocks base method.
func (m *MockInputPath) Path() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

// Path indicates an expected call of Path.
func (mr *MockInputPathMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockInputPath)(nil).Path))
}
