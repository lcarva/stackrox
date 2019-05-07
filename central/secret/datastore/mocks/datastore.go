// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/secret/datastore (interfaces: DataStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
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

// CountSecrets mocks base method
func (m *MockDataStore) CountSecrets(arg0 context.Context) (int, error) {
	ret := m.ctrl.Call(m, "CountSecrets", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountSecrets indicates an expected call of CountSecrets
func (mr *MockDataStoreMockRecorder) CountSecrets(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountSecrets", reflect.TypeOf((*MockDataStore)(nil).CountSecrets), arg0)
}

// GetSecret mocks base method
func (m *MockDataStore) GetSecret(arg0 context.Context, arg1 string) (*storage.Secret, bool, error) {
	ret := m.ctrl.Call(m, "GetSecret", arg0, arg1)
	ret0, _ := ret[0].(*storage.Secret)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSecret indicates an expected call of GetSecret
func (mr *MockDataStoreMockRecorder) GetSecret(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockDataStore)(nil).GetSecret), arg0, arg1)
}

// ListSecrets mocks base method
func (m *MockDataStore) ListSecrets(arg0 context.Context) ([]*storage.ListSecret, error) {
	ret := m.ctrl.Call(m, "ListSecrets", arg0)
	ret0, _ := ret[0].([]*storage.ListSecret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets
func (mr *MockDataStoreMockRecorder) ListSecrets(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockDataStore)(nil).ListSecrets), arg0)
}

// RemoveSecret mocks base method
func (m *MockDataStore) RemoveSecret(arg0 context.Context, arg1 string) error {
	ret := m.ctrl.Call(m, "RemoveSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSecret indicates an expected call of RemoveSecret
func (mr *MockDataStoreMockRecorder) RemoveSecret(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSecret", reflect.TypeOf((*MockDataStore)(nil).RemoveSecret), arg0, arg1)
}

// Search mocks base method
func (m *MockDataStore) Search(arg0 context.Context, arg1 *v1.Query) ([]search.Result, error) {
	ret := m.ctrl.Call(m, "Search", arg0, arg1)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockDataStoreMockRecorder) Search(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDataStore)(nil).Search), arg0, arg1)
}

// SearchListSecrets mocks base method
func (m *MockDataStore) SearchListSecrets(arg0 context.Context, arg1 *v1.Query) ([]*storage.ListSecret, error) {
	ret := m.ctrl.Call(m, "SearchListSecrets", arg0, arg1)
	ret0, _ := ret[0].([]*storage.ListSecret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchListSecrets indicates an expected call of SearchListSecrets
func (mr *MockDataStoreMockRecorder) SearchListSecrets(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchListSecrets", reflect.TypeOf((*MockDataStore)(nil).SearchListSecrets), arg0, arg1)
}

// SearchSecrets mocks base method
func (m *MockDataStore) SearchSecrets(arg0 context.Context, arg1 *v1.Query) ([]*v1.SearchResult, error) {
	ret := m.ctrl.Call(m, "SearchSecrets", arg0, arg1)
	ret0, _ := ret[0].([]*v1.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchSecrets indicates an expected call of SearchSecrets
func (mr *MockDataStoreMockRecorder) SearchSecrets(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchSecrets", reflect.TypeOf((*MockDataStore)(nil).SearchSecrets), arg0, arg1)
}

// UpsertSecret mocks base method
func (m *MockDataStore) UpsertSecret(arg0 context.Context, arg1 *storage.Secret) error {
	ret := m.ctrl.Call(m, "UpsertSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertSecret indicates an expected call of UpsertSecret
func (mr *MockDataStoreMockRecorder) UpsertSecret(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertSecret", reflect.TypeOf((*MockDataStore)(nil).UpsertSecret), arg0, arg1)
}
