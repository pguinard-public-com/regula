// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fugue/regula/v3/pkg/loader (interfaces: IACConfiguration)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	loader "github.com/fugue/regula/v3/pkg/loader"
	gomock "github.com/golang/mock/gomock"
)

// MockIACConfiguration is a mock of IACConfiguration interface.
type MockIACConfiguration struct {
	ctrl     *gomock.Controller
	recorder *MockIACConfigurationMockRecorder
}

// MockIACConfigurationMockRecorder is the mock recorder for MockIACConfiguration.
type MockIACConfigurationMockRecorder struct {
	mock *MockIACConfiguration
}

// NewMockIACConfiguration creates a new mock instance.
func NewMockIACConfiguration(ctrl *gomock.Controller) *MockIACConfiguration {
	mock := &MockIACConfiguration{ctrl: ctrl}
	mock.recorder = &MockIACConfigurationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIACConfiguration) EXPECT() *MockIACConfigurationMockRecorder {
	return m.recorder
}

// LoadedFiles mocks base method.
func (m *MockIACConfiguration) LoadedFiles() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadedFiles")
	ret0, _ := ret[0].([]string)
	return ret0
}

// LoadedFiles indicates an expected call of LoadedFiles.
func (mr *MockIACConfigurationMockRecorder) LoadedFiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadedFiles", reflect.TypeOf((*MockIACConfiguration)(nil).LoadedFiles))
}

// Location mocks base method.
func (m *MockIACConfiguration) Location(arg0 []string) ([]loader.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location", arg0)
	ret0, _ := ret[0].([]loader.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Location indicates an expected call of Location.
func (mr *MockIACConfigurationMockRecorder) Location(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockIACConfiguration)(nil).Location), arg0)
}

// RegulaInput mocks base method.
func (m *MockIACConfiguration) RegulaInput() loader.RegulaInput {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegulaInput")
	ret0, _ := ret[0].(loader.RegulaInput)
	return ret0
}

// RegulaInput indicates an expected call of RegulaInput.
func (mr *MockIACConfigurationMockRecorder) RegulaInput() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegulaInput", reflect.TypeOf((*MockIACConfiguration)(nil).RegulaInput))
}
