// Code generated by MockGen. DO NOT EDIT.
// Source: shared/storage.go

// Package shared is a generated GoMock package.
package shared

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockReadable is a mock of Readable interface
type MockReadable struct {
	ctrl     *gomock.Controller
	recorder *MockReadableMockRecorder
}

// MockReadableMockRecorder is the mock recorder for MockReadable
type MockReadableMockRecorder struct {
	mock *MockReadable
}

// NewMockReadable creates a new mock instance
func NewMockReadable(ctrl *gomock.Controller) *MockReadable {
	mock := &MockReadable{ctrl: ctrl}
	mock.recorder = &MockReadableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReadable) EXPECT() *MockReadableMockRecorder {
	return m.recorder
}

// NewReadCloser mocks base method
func (m *MockReadable) NewReadCloser(arg0 interface{}) (io.ReadCloser, error) {
	ret := m.ctrl.Call(m, "NewReadCloser", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewReadCloser indicates an expected call of NewReadCloser
func (mr *MockReadableMockRecorder) NewReadCloser(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReadCloser", reflect.TypeOf((*MockReadable)(nil).NewReadCloser), arg0)
}

// MockReadWritable is a mock of ReadWritable interface
type MockReadWritable struct {
	ctrl     *gomock.Controller
	recorder *MockReadWritableMockRecorder
}

// MockReadWritableMockRecorder is the mock recorder for MockReadWritable
type MockReadWritableMockRecorder struct {
	mock *MockReadWritable
}

// NewMockReadWritable creates a new mock instance
func NewMockReadWritable(ctrl *gomock.Controller) *MockReadWritable {
	mock := &MockReadWritable{ctrl: ctrl}
	mock.recorder = &MockReadWritableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReadWritable) EXPECT() *MockReadWritableMockRecorder {
	return m.recorder
}

// NewReadCloser mocks base method
func (m *MockReadWritable) NewReadCloser(arg0 interface{}) (io.ReadCloser, error) {
	ret := m.ctrl.Call(m, "NewReadCloser", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewReadCloser indicates an expected call of NewReadCloser
func (mr *MockReadWritableMockRecorder) NewReadCloser(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReadCloser", reflect.TypeOf((*MockReadWritable)(nil).NewReadCloser), arg0)
}

// NewWriteCloser mocks base method
func (m *MockReadWritable) NewWriteCloser(arg0 interface{}) (io.WriteCloser, error) {
	ret := m.ctrl.Call(m, "NewWriteCloser", arg0)
	ret0, _ := ret[0].(io.WriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewWriteCloser indicates an expected call of NewWriteCloser
func (mr *MockReadWritableMockRecorder) NewWriteCloser(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWriteCloser", reflect.TypeOf((*MockReadWritable)(nil).NewWriteCloser), arg0)
}

// MockCachedStore is a mock of CachedStore interface
type MockCachedStore struct {
	ctrl     *gomock.Controller
	recorder *MockCachedStoreMockRecorder
}

// MockCachedStoreMockRecorder is the mock recorder for MockCachedStore
type MockCachedStoreMockRecorder struct {
	mock *MockCachedStore
}

// NewMockCachedStore creates a new mock instance
func NewMockCachedStore(ctrl *gomock.Controller) *MockCachedStore {
	mock := &MockCachedStore{ctrl: ctrl}
	mock.recorder = &MockCachedStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCachedStore) EXPECT() *MockCachedStoreMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockCachedStore) Get(cacheID, storeID, value interface{}) error {
	ret := m.ctrl.Call(m, "Get", cacheID, storeID, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockCachedStoreMockRecorder) Get(cacheID, storeID, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCachedStore)(nil).Get), cacheID, storeID, value)
}
