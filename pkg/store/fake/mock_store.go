// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/octant/pkg/store (interfaces: Store)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"

	cluster "github.com/vmware-tanzu/octant/internal/cluster"
	store "github.com/vmware-tanzu/octant/pkg/store"
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

// Create mocks base method
func (m *MockStore) Create(arg0 context.Context, arg1 *unstructured.Unstructured) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockStore) Delete(arg0 context.Context, arg1 store.Key) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStore)(nil).Delete), arg0, arg1)
}

// Get mocks base method
func (m *MockStore) Get(arg0 context.Context, arg1 store.Key) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), arg0, arg1)
}

// IsLoading mocks base method
func (m *MockStore) IsLoading(arg0 context.Context, arg1 store.Key) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLoading", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLoading indicates an expected call of IsLoading
func (mr *MockStoreMockRecorder) IsLoading(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLoading", reflect.TypeOf((*MockStore)(nil).IsLoading), arg0, arg1)
}

// List mocks base method
func (m *MockStore) List(arg0 context.Context, arg1 store.Key) (*unstructured.UnstructuredList, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*unstructured.UnstructuredList)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List
func (mr *MockStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStore)(nil).List), arg0, arg1)
}

// RegisterOnUpdate mocks base method
func (m *MockStore) RegisterOnUpdate(arg0 store.UpdateFn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterOnUpdate", arg0)
}

// RegisterOnUpdate indicates an expected call of RegisterOnUpdate
func (mr *MockStoreMockRecorder) RegisterOnUpdate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterOnUpdate", reflect.TypeOf((*MockStore)(nil).RegisterOnUpdate), arg0)
}

// Unwatch mocks base method
func (m *MockStore) Unwatch(arg0 context.Context, arg1 ...schema.GroupVersionKind) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Unwatch", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unwatch indicates an expected call of Unwatch
func (mr *MockStoreMockRecorder) Unwatch(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unwatch", reflect.TypeOf((*MockStore)(nil).Unwatch), varargs...)
}

// Update mocks base method
func (m *MockStore) Update(arg0 context.Context, arg1 store.Key, arg2 func(*unstructured.Unstructured) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockStoreMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStore)(nil).Update), arg0, arg1, arg2)
}

// UpdateClusterClient mocks base method
func (m *MockStore) UpdateClusterClient(arg0 context.Context, arg1 cluster.ClientInterface) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClusterClient", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClusterClient indicates an expected call of UpdateClusterClient
func (mr *MockStoreMockRecorder) UpdateClusterClient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClusterClient", reflect.TypeOf((*MockStore)(nil).UpdateClusterClient), arg0, arg1)
}

// Watch mocks base method
func (m *MockStore) Watch(arg0 context.Context, arg1 store.Key, arg2 cache.ResourceEventHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Watch indicates an expected call of Watch
func (mr *MockStoreMockRecorder) Watch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockStore)(nil).Watch), arg0, arg1, arg2)
}
