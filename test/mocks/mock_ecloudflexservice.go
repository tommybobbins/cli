// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ukfast/sdk-go/pkg/service/ecloudflex (interfaces: ECloudFlexService)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	connection "github.com/ukfast/sdk-go/pkg/connection"
	ecloudflex "github.com/ukfast/sdk-go/pkg/service/ecloudflex"
	reflect "reflect"
)

// MockECloudFlexService is a mock of ECloudFlexService interface
type MockECloudFlexService struct {
	ctrl     *gomock.Controller
	recorder *MockECloudFlexServiceMockRecorder
}

// MockECloudFlexServiceMockRecorder is the mock recorder for MockECloudFlexService
type MockECloudFlexServiceMockRecorder struct {
	mock *MockECloudFlexService
}

// NewMockECloudFlexService creates a new mock instance
func NewMockECloudFlexService(ctrl *gomock.Controller) *MockECloudFlexService {
	mock := &MockECloudFlexService{ctrl: ctrl}
	mock.recorder = &MockECloudFlexServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockECloudFlexService) EXPECT() *MockECloudFlexServiceMockRecorder {
	return m.recorder
}

// GetProject mocks base method
func (m *MockECloudFlexService) GetProject(arg0 int) (ecloudflex.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", arg0)
	ret0, _ := ret[0].(ecloudflex.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject
func (mr *MockECloudFlexServiceMockRecorder) GetProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockECloudFlexService)(nil).GetProject), arg0)
}

// GetProjects mocks base method
func (m *MockECloudFlexService) GetProjects(arg0 connection.APIRequestParameters) ([]ecloudflex.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjects", arg0)
	ret0, _ := ret[0].([]ecloudflex.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjects indicates an expected call of GetProjects
func (mr *MockECloudFlexServiceMockRecorder) GetProjects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjects", reflect.TypeOf((*MockECloudFlexService)(nil).GetProjects), arg0)
}

// GetProjectsPaginated mocks base method
func (m *MockECloudFlexService) GetProjectsPaginated(arg0 connection.APIRequestParameters) (*ecloudflex.PaginatedProject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectsPaginated", arg0)
	ret0, _ := ret[0].(*ecloudflex.PaginatedProject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectsPaginated indicates an expected call of GetProjectsPaginated
func (mr *MockECloudFlexServiceMockRecorder) GetProjectsPaginated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectsPaginated", reflect.TypeOf((*MockECloudFlexService)(nil).GetProjectsPaginated), arg0)
}