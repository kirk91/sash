// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package config is a generated GoMock package.
package config

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockStore) Get(namespace, typ, key string) ([]byte, error) {
	ret := m.ctrl.Call(m, "Get", namespace, typ, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockStoreMockRecorder) Get(namespace, typ, key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), namespace, typ, key)
}

// Add mocks base method
func (m *MockStore) Add(namespace, typ, key string, value []byte) error {
	ret := m.ctrl.Call(m, "Add", namespace, typ, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockStoreMockRecorder) Add(namespace, typ, key, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStore)(nil).Add), namespace, typ, key, value)
}

// Update mocks base method
func (m *MockStore) Update(namespace, typ, key string, value []byte) error {
	ret := m.ctrl.Call(m, "Update", namespace, typ, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockStoreMockRecorder) Update(namespace, typ, key, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStore)(nil).Update), namespace, typ, key, value)
}

// Del mocks base method
func (m *MockStore) Del(namespace, typ, key string) error {
	ret := m.ctrl.Call(m, "Del", namespace, typ, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del
func (mr *MockStoreMockRecorder) Del(namespace, typ, key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockStore)(nil).Del), namespace, typ, key)
}

// Exist mocks base method
func (m *MockStore) Exist(namespace, typ, key string) bool {
	ret := m.ctrl.Call(m, "Exist", namespace, typ, key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exist indicates an expected call of Exist
func (mr *MockStoreMockRecorder) Exist(namespace, typ, key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockStore)(nil).Exist), namespace, typ, key)
}

// GetKeys mocks base method
func (m *MockStore) GetKeys(namespace, typ string) ([]string, error) {
	ret := m.ctrl.Call(m, "GetKeys", namespace, typ)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeys indicates an expected call of GetKeys
func (mr *MockStoreMockRecorder) GetKeys(namespace, typ interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeys", reflect.TypeOf((*MockStore)(nil).GetKeys), namespace, typ)
}

// Start mocks base method
func (m *MockStore) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockStoreMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockStore)(nil).Start))
}

// Stop mocks base method
func (m *MockStore) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockStoreMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockStore)(nil).Stop))
}

// MockSubscribableStore is a mock of SubscribableStore interface
type MockSubscribableStore struct {
	ctrl     *gomock.Controller
	recorder *MockSubscribableStoreMockRecorder
}

// MockSubscribableStoreMockRecorder is the mock recorder for MockSubscribableStore
type MockSubscribableStoreMockRecorder struct {
	mock *MockSubscribableStore
}

// NewMockSubscribableStore creates a new mock instance
func NewMockSubscribableStore(ctrl *gomock.Controller) *MockSubscribableStore {
	mock := &MockSubscribableStore{ctrl: ctrl}
	mock.recorder = &MockSubscribableStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSubscribableStore) EXPECT() *MockSubscribableStoreMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockSubscribableStore) Get(namespace, typ, key string) ([]byte, error) {
	ret := m.ctrl.Call(m, "Get", namespace, typ, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSubscribableStoreMockRecorder) Get(namespace, typ, key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSubscribableStore)(nil).Get), namespace, typ, key)
}

// Add mocks base method
func (m *MockSubscribableStore) Add(namespace, typ, key string, value []byte) error {
	ret := m.ctrl.Call(m, "Add", namespace, typ, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockSubscribableStoreMockRecorder) Add(namespace, typ, key, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockSubscribableStore)(nil).Add), namespace, typ, key, value)
}

// Update mocks base method
func (m *MockSubscribableStore) Update(namespace, typ, key string, value []byte) error {
	ret := m.ctrl.Call(m, "Update", namespace, typ, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockSubscribableStoreMockRecorder) Update(namespace, typ, key, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSubscribableStore)(nil).Update), namespace, typ, key, value)
}

// Del mocks base method
func (m *MockSubscribableStore) Del(namespace, typ, key string) error {
	ret := m.ctrl.Call(m, "Del", namespace, typ, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del
func (mr *MockSubscribableStoreMockRecorder) Del(namespace, typ, key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockSubscribableStore)(nil).Del), namespace, typ, key)
}

// Exist mocks base method
func (m *MockSubscribableStore) Exist(namespace, typ, key string) bool {
	ret := m.ctrl.Call(m, "Exist", namespace, typ, key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exist indicates an expected call of Exist
func (mr *MockSubscribableStoreMockRecorder) Exist(namespace, typ, key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockSubscribableStore)(nil).Exist), namespace, typ, key)
}

// GetKeys mocks base method
func (m *MockSubscribableStore) GetKeys(namespace, typ string) ([]string, error) {
	ret := m.ctrl.Call(m, "GetKeys", namespace, typ)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeys indicates an expected call of GetKeys
func (mr *MockSubscribableStoreMockRecorder) GetKeys(namespace, typ interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeys", reflect.TypeOf((*MockSubscribableStore)(nil).GetKeys), namespace, typ)
}

// Start mocks base method
func (m *MockSubscribableStore) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockSubscribableStoreMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockSubscribableStore)(nil).Start))
}

// Stop mocks base method
func (m *MockSubscribableStore) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockSubscribableStoreMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockSubscribableStore)(nil).Stop))
}

// Subscribe mocks base method
func (m *MockSubscribableStore) Subscribe(namespace string) error {
	ret := m.ctrl.Call(m, "Subscribe", namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockSubscribableStoreMockRecorder) Subscribe(namespace interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockSubscribableStore)(nil).Subscribe), namespace)
}

// UnSubscribe mocks base method
func (m *MockSubscribableStore) UnSubscribe(namespace string) error {
	ret := m.ctrl.Call(m, "UnSubscribe", namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnSubscribe indicates an expected call of UnSubscribe
func (mr *MockSubscribableStoreMockRecorder) UnSubscribe(namespace interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnSubscribe", reflect.TypeOf((*MockSubscribableStore)(nil).UnSubscribe), namespace)
}

// Event mocks base method
func (m *MockSubscribableStore) Event() <-chan struct{} {
	ret := m.ctrl.Call(m, "Event")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Event indicates an expected call of Event
func (mr *MockSubscribableStoreMockRecorder) Event() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Event", reflect.TypeOf((*MockSubscribableStore)(nil).Event))
}
