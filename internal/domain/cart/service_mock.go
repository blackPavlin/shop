// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package cart is a generated GoMock package.
package cart

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Query mocks base method.
func (m *MockService) Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", ctx, criteria)
	ret0, _ := ret[0].(*QueryResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockServiceMockRecorder) Query(ctx, criteria interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockService)(nil).Query), ctx, criteria)
}

// Save mocks base method.
func (m *MockService) Save(ctx context.Context, props *Props) (*Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, props)
	ret0, _ := ret[0].(*Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockServiceMockRecorder) Save(ctx, props interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockService)(nil).Save), ctx, props)
}
