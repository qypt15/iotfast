// Code generated by MockGen. DO NOT EDIT.
// Source: server/plugin.go

// Package server is a generated GoMock package.
package server

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPlugin is a mock of Plugin interface
type MockPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockPluginMockRecorder
}

// MockPluginMockRecorder is the mock recorder for MockPlugin
type MockPluginMockRecorder struct {
	mock *MockPlugin
}

// NewMockPlugin creates a new mock instance
func NewMockPlugin(ctrl *gomock.Controller) *MockPlugin {
	mock := &MockPlugin{ctrl: ctrl}
	mock.recorder = &MockPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlugin) EXPECT() *MockPluginMockRecorder {
	return m.recorder
}

// Load mocks base method
func (m *MockPlugin) Load(service Server) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", service)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load
func (mr *MockPluginMockRecorder) Load(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockPlugin)(nil).Load), service)
}

// Unload mocks base method
func (m *MockPlugin) Unload() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unload")
	ret0, _ := ret[0].(error)
	return ret0
}

// Unload indicates an expected call of Unload
func (mr *MockPluginMockRecorder) Unload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unload", reflect.TypeOf((*MockPlugin)(nil).Unload))
}

// HookWrapper mocks base method
func (m *MockPlugin) HookWrapper() HookWrapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HookWrapper")
	ret0, _ := ret[0].(HookWrapper)
	return ret0
}

// HookWrapper indicates an expected call of HookWrapper
func (mr *MockPluginMockRecorder) HookWrapper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HookWrapper", reflect.TypeOf((*MockPlugin)(nil).HookWrapper))
}

// Name mocks base method
func (m *MockPlugin) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockPluginMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockPlugin)(nil).Name))
}