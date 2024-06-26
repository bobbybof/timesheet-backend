// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bobbybof/timesheet/repository (interfaces: Store)
//
// Generated by this command:
//
//	mockgen -package mock -destination repository/mock/store.go github.com/bobbybof/timesheet/repository Store
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	repository "github.com/bobbybof/timesheet/internal/repository"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateActivity mocks base method.
func (m *MockStore) CreateActivity(arg0 context.Context, arg1 repository.CreateActivityParams) (repository.Activity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateActivity", arg0, arg1)
	ret0, _ := ret[0].(repository.Activity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateActivity indicates an expected call of CreateActivity.
func (mr *MockStoreMockRecorder) CreateActivity(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateActivity", reflect.TypeOf((*MockStore)(nil).CreateActivity), arg0, arg1)
}

// CreateProject mocks base method.
func (m *MockStore) CreateProject(arg0 context.Context, arg1 string) (repository.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0, arg1)
	ret0, _ := ret[0].(repository.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockStoreMockRecorder) CreateProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockStore)(nil).CreateProject), arg0, arg1)
}

// DeleteActivity mocks base method.
func (m *MockStore) DeleteActivity(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteActivity", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteActivity indicates an expected call of DeleteActivity.
func (mr *MockStoreMockRecorder) DeleteActivity(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteActivity", reflect.TypeOf((*MockStore)(nil).DeleteActivity), arg0, arg1)
}

// GetProjects mocks base method.
func (m *MockStore) GetProjects(arg0 context.Context) ([]repository.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjects", arg0)
	ret0, _ := ret[0].([]repository.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjects indicates an expected call of GetProjects.
func (mr *MockStoreMockRecorder) GetProjects(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjects", reflect.TypeOf((*MockStore)(nil).GetProjects), arg0)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 int32) (repository.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(repository.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// UpdateActivity mocks base method.
func (m *MockStore) UpdateActivity(arg0 context.Context, arg1 repository.UpdateActivityParams) (repository.Activity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateActivity", arg0, arg1)
	ret0, _ := ret[0].(repository.Activity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateActivity indicates an expected call of UpdateActivity.
func (mr *MockStoreMockRecorder) UpdateActivity(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateActivity", reflect.TypeOf((*MockStore)(nil).UpdateActivity), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 repository.UpdateUserParams) (repository.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(repository.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
