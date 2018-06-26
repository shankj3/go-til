// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_consul is a generated GoMock package.
package mock_consul

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/hashicorp/consul/api"
	reflect "reflect"
)

// MockConsuletty is a mock of Consuletty interface
type MockConsuletty struct {
	ctrl     *gomock.Controller
	recorder *MockConsulettyMockRecorder
}

// MockConsulettyMockRecorder is the mock recorder for MockConsuletty
type MockConsulettyMockRecorder struct {
	mock *MockConsuletty
}

// NewMockConsuletty creates a new mock instance
func NewMockConsuletty(ctrl *gomock.Controller) *MockConsuletty {
	mock := &MockConsuletty{ctrl: ctrl}
	mock.recorder = &MockConsulettyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConsuletty) EXPECT() *MockConsulettyMockRecorder {
	return m.recorder
}

// AddKeyValue mocks base method
func (m *MockConsuletty) AddKeyValue(key string, value []byte) error {
	ret := m.ctrl.Call(m, "AddKeyValue", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddKeyValue indicates an expected call of AddKeyValue
func (mr *MockConsulettyMockRecorder) AddKeyValue(key, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddKeyValue", reflect.TypeOf((*MockConsuletty)(nil).AddKeyValue), key, value)
}

// RemoveValue mocks base method
func (m *MockConsuletty) RemoveValue(key string) error {
	ret := m.ctrl.Call(m, "RemoveValue", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveValue indicates an expected call of RemoveValue
func (mr *MockConsulettyMockRecorder) RemoveValue(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveValue", reflect.TypeOf((*MockConsuletty)(nil).RemoveValue), key)
}

// RemoveValues mocks base method
func (m *MockConsuletty) RemoveValues(prefix string) error {
	ret := m.ctrl.Call(m, "RemoveValues", prefix)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveValues indicates an expected call of RemoveValues
func (mr *MockConsulettyMockRecorder) RemoveValues(prefix interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveValues", reflect.TypeOf((*MockConsuletty)(nil).RemoveValues), prefix)
}

// GetKeys mocks base method
func (m *MockConsuletty) GetKeys(prefix string) ([]string, error) {
	ret := m.ctrl.Call(m, "GetKeys", prefix)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeys indicates an expected call of GetKeys
func (mr *MockConsulettyMockRecorder) GetKeys(prefix interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeys", reflect.TypeOf((*MockConsuletty)(nil).GetKeys), prefix)
}

// GetKeyValue mocks base method
func (m *MockConsuletty) GetKeyValue(key string) (*api.KVPair, error) {
	ret := m.ctrl.Call(m, "GetKeyValue", key)
	ret0, _ := ret[0].(*api.KVPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyValue indicates an expected call of GetKeyValue
func (mr *MockConsulettyMockRecorder) GetKeyValue(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyValue", reflect.TypeOf((*MockConsuletty)(nil).GetKeyValue), key)
}

// GetKeyValues mocks base method
func (m *MockConsuletty) GetKeyValues(prefix string) (api.KVPairs, error) {
	ret := m.ctrl.Call(m, "GetKeyValues", prefix)
	ret0, _ := ret[0].(api.KVPairs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyValues indicates an expected call of GetKeyValues
func (mr *MockConsulettyMockRecorder) GetKeyValues(prefix interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyValues", reflect.TypeOf((*MockConsuletty)(nil).GetKeyValues), prefix)
}

// RegisterService mocks base method
func (m *MockConsuletty) RegisterService(addr string, port int, name string) error {
	ret := m.ctrl.Call(m, "RegisterService", addr, port, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterService indicates an expected call of RegisterService
func (mr *MockConsulettyMockRecorder) RegisterService(addr, port, name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterService", reflect.TypeOf((*MockConsuletty)(nil).RegisterService), addr, port, name)
}

// RemoveService mocks base method
func (m *MockConsuletty) RemoveService(name string) error {
	ret := m.ctrl.Call(m, "RemoveService", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveService indicates an expected call of RemoveService
func (mr *MockConsulettyMockRecorder) RemoveService(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveService", reflect.TypeOf((*MockConsuletty)(nil).RemoveService), name)
}

// IsConnected mocks base method
func (m *MockConsuletty) IsConnected() bool {
	ret := m.ctrl.Call(m, "IsConnected")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsConnected indicates an expected call of IsConnected
func (mr *MockConsulettyMockRecorder) IsConnected() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConnected", reflect.TypeOf((*MockConsuletty)(nil).IsConnected))
}
