// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/processwhitelist/datastore (interfaces: DataStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	reflect "reflect"
)

// MockDataStore is a mock of DataStore interface
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// AddProcessWhitelist mocks base method
func (m *MockDataStore) AddProcessWhitelist(arg0 context.Context, arg1 *storage.ProcessWhitelist) (string, error) {
	ret := m.ctrl.Call(m, "AddProcessWhitelist", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProcessWhitelist indicates an expected call of AddProcessWhitelist
func (mr *MockDataStoreMockRecorder) AddProcessWhitelist(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProcessWhitelist", reflect.TypeOf((*MockDataStore)(nil).AddProcessWhitelist), arg0, arg1)
}

// GetProcessWhitelist mocks base method
func (m *MockDataStore) GetProcessWhitelist(arg0 context.Context, arg1 *storage.ProcessWhitelistKey) (*storage.ProcessWhitelist, error) {
	ret := m.ctrl.Call(m, "GetProcessWhitelist", arg0, arg1)
	ret0, _ := ret[0].(*storage.ProcessWhitelist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProcessWhitelist indicates an expected call of GetProcessWhitelist
func (mr *MockDataStoreMockRecorder) GetProcessWhitelist(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProcessWhitelist", reflect.TypeOf((*MockDataStore)(nil).GetProcessWhitelist), arg0, arg1)
}

// GetProcessWhitelists mocks base method
func (m *MockDataStore) GetProcessWhitelists(arg0 context.Context) ([]*storage.ProcessWhitelist, error) {
	ret := m.ctrl.Call(m, "GetProcessWhitelists", arg0)
	ret0, _ := ret[0].([]*storage.ProcessWhitelist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProcessWhitelists indicates an expected call of GetProcessWhitelists
func (mr *MockDataStoreMockRecorder) GetProcessWhitelists(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProcessWhitelists", reflect.TypeOf((*MockDataStore)(nil).GetProcessWhitelists), arg0)
}

// RemoveProcessWhitelist mocks base method
func (m *MockDataStore) RemoveProcessWhitelist(arg0 context.Context, arg1 *storage.ProcessWhitelistKey) error {
	ret := m.ctrl.Call(m, "RemoveProcessWhitelist", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProcessWhitelist indicates an expected call of RemoveProcessWhitelist
func (mr *MockDataStoreMockRecorder) RemoveProcessWhitelist(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProcessWhitelist", reflect.TypeOf((*MockDataStore)(nil).RemoveProcessWhitelist), arg0, arg1)
}

// RoxLockProcessWhitelist mocks base method
func (m *MockDataStore) RoxLockProcessWhitelist(arg0 context.Context, arg1 *storage.ProcessWhitelistKey, arg2 bool) (*storage.ProcessWhitelist, error) {
	ret := m.ctrl.Call(m, "RoxLockProcessWhitelist", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.ProcessWhitelist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoxLockProcessWhitelist indicates an expected call of RoxLockProcessWhitelist
func (mr *MockDataStoreMockRecorder) RoxLockProcessWhitelist(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoxLockProcessWhitelist", reflect.TypeOf((*MockDataStore)(nil).RoxLockProcessWhitelist), arg0, arg1, arg2)
}

// SearchRawProcessWhitelists mocks base method
func (m *MockDataStore) SearchRawProcessWhitelists(arg0 context.Context, arg1 *v1.Query) ([]*storage.ProcessWhitelist, error) {
	ret := m.ctrl.Call(m, "SearchRawProcessWhitelists", arg0, arg1)
	ret0, _ := ret[0].([]*storage.ProcessWhitelist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRawProcessWhitelists indicates an expected call of SearchRawProcessWhitelists
func (mr *MockDataStoreMockRecorder) SearchRawProcessWhitelists(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRawProcessWhitelists", reflect.TypeOf((*MockDataStore)(nil).SearchRawProcessWhitelists), arg0, arg1)
}

// UpdateProcessWhitelistElements mocks base method
func (m *MockDataStore) UpdateProcessWhitelistElements(arg0 context.Context, arg1 *storage.ProcessWhitelistKey, arg2, arg3 []*storage.WhitelistItem, arg4 bool) (*storage.ProcessWhitelist, error) {
	ret := m.ctrl.Call(m, "UpdateProcessWhitelistElements", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*storage.ProcessWhitelist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProcessWhitelistElements indicates an expected call of UpdateProcessWhitelistElements
func (mr *MockDataStoreMockRecorder) UpdateProcessWhitelistElements(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProcessWhitelistElements", reflect.TypeOf((*MockDataStore)(nil).UpdateProcessWhitelistElements), arg0, arg1, arg2, arg3, arg4)
}

// UpsertProcessWhitelist mocks base method
func (m *MockDataStore) UpsertProcessWhitelist(arg0 context.Context, arg1 *storage.ProcessWhitelistKey, arg2 []*storage.WhitelistItem, arg3 bool) (*storage.ProcessWhitelist, error) {
	ret := m.ctrl.Call(m, "UpsertProcessWhitelist", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*storage.ProcessWhitelist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertProcessWhitelist indicates an expected call of UpsertProcessWhitelist
func (mr *MockDataStoreMockRecorder) UpsertProcessWhitelist(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertProcessWhitelist", reflect.TypeOf((*MockDataStore)(nil).UpsertProcessWhitelist), arg0, arg1, arg2, arg3)
}

// UserLockProcessWhitelist mocks base method
func (m *MockDataStore) UserLockProcessWhitelist(arg0 context.Context, arg1 *storage.ProcessWhitelistKey, arg2 bool) (*storage.ProcessWhitelist, error) {
	ret := m.ctrl.Call(m, "UserLockProcessWhitelist", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.ProcessWhitelist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLockProcessWhitelist indicates an expected call of UserLockProcessWhitelist
func (mr *MockDataStoreMockRecorder) UserLockProcessWhitelist(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLockProcessWhitelist", reflect.TypeOf((*MockDataStore)(nil).UserLockProcessWhitelist), arg0, arg1, arg2)
}
