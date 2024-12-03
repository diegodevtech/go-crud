// Code generated by MockGen. DO NOT EDIT.
// Source: src/model/service/user_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	rest_err "github.com/diegodevtech/go-crud/src/configuration/rest_err"
	model "github.com/diegodevtech/go-crud/src/model"
	gomock "github.com/golang/mock/gomock"
)

// MockUserDomainService is a mock of UserDomainService interface.
type MockUserDomainService struct {
	ctrl     *gomock.Controller
	recorder *MockUserDomainServiceMockRecorder
}

// MockUserDomainServiceMockRecorder is the mock recorder for MockUserDomainService.
type MockUserDomainServiceMockRecorder struct {
	mock *MockUserDomainService
}

// NewMockUserDomainService creates a new mock instance.
func NewMockUserDomainService(ctrl *gomock.Controller) *MockUserDomainService {
	mock := &MockUserDomainService{ctrl: ctrl}
	mock.recorder = &MockUserDomainServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDomainService) EXPECT() *MockUserDomainServiceMockRecorder {
	return m.recorder
}

// CreateUserService mocks base method.
func (m *MockUserDomainService) CreateUserService(arg0 model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserService", arg0)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// CreateUserService indicates an expected call of CreateUserService.
func (mr *MockUserDomainServiceMockRecorder) CreateUserService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserService", reflect.TypeOf((*MockUserDomainService)(nil).CreateUserService), arg0)
}

// DeleteUserService mocks base method.
func (m *MockUserDomainService) DeleteUserService(arg0 string) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserService", arg0)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

// DeleteUserService indicates an expected call of DeleteUserService.
func (mr *MockUserDomainServiceMockRecorder) DeleteUserService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserService", reflect.TypeOf((*MockUserDomainService)(nil).DeleteUserService), arg0)
}

// FindUserByEmailService mocks base method.
func (m *MockUserDomainService) FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmailService", email)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// FindUserByEmailService indicates an expected call of FindUserByEmailService.
func (mr *MockUserDomainServiceMockRecorder) FindUserByEmailService(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmailService", reflect.TypeOf((*MockUserDomainService)(nil).FindUserByEmailService), email)
}

// FindUserByIDService mocks base method.
func (m *MockUserDomainService) FindUserByIDService(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByIDService", id)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// FindUserByIDService indicates an expected call of FindUserByIDService.
func (mr *MockUserDomainServiceMockRecorder) FindUserByIDService(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByIDService", reflect.TypeOf((*MockUserDomainService)(nil).FindUserByIDService), id)
}

// LoginService mocks base method.
func (m *MockUserDomainService) LoginService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginService", userDomain)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(*rest_err.RestErr)
	return ret0, ret1, ret2
}

// LoginService indicates an expected call of LoginService.
func (mr *MockUserDomainServiceMockRecorder) LoginService(userDomain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginService", reflect.TypeOf((*MockUserDomainService)(nil).LoginService), userDomain)
}

// UpdateUserService mocks base method.
func (m *MockUserDomainService) UpdateUserService(arg0 string, arg1 model.UserDomainInterface) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserService", arg0, arg1)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

// UpdateUserService indicates an expected call of UpdateUserService.
func (mr *MockUserDomainServiceMockRecorder) UpdateUserService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserService", reflect.TypeOf((*MockUserDomainService)(nil).UpdateUserService), arg0, arg1)
}
