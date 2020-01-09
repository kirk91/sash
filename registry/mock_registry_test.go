// Code generated by MockGen. DO NOT EDIT.
// Source: ../model/service.go

// Package registry is a generated GoMock package.
package registry

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/samaritan-proxy/sash/model"
	reflect "reflect"
)

// MockServiceRegistry is a mock of ServiceRegistry interface
type MockServiceRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockServiceRegistryMockRecorder
}

// MockServiceRegistryMockRecorder is the mock recorder for MockServiceRegistry
type MockServiceRegistryMockRecorder struct {
	mock *MockServiceRegistry
}

// NewMockServiceRegistry creates a new mock instance
func NewMockServiceRegistry(ctrl *gomock.Controller) *MockServiceRegistry {
	mock := &MockServiceRegistry{ctrl: ctrl}
	mock.recorder = &MockServiceRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceRegistry) EXPECT() *MockServiceRegistryMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockServiceRegistry) Run(ctx context.Context) {
	m.ctrl.Call(m, "Run", ctx)
}

// Run indicates an expected call of Run
func (mr *MockServiceRegistryMockRecorder) Run(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockServiceRegistry)(nil).Run), ctx)
}

// List mocks base method
func (m *MockServiceRegistry) List() ([]string, error) {
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockServiceRegistryMockRecorder) List() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServiceRegistry)(nil).List))
}

// Get mocks base method
func (m *MockServiceRegistry) Get(name string) (*model.Service, error) {
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(*model.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockServiceRegistryMockRecorder) Get(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServiceRegistry)(nil).Get), name)
}
