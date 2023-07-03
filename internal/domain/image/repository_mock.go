// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package image is a generated GoMock package.
package image

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// BulkCreateTx mocks base method.
func (m *MockRepository) BulkCreateTx(ctx context.Context, images Images) (Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateTx", ctx, images)
	ret0, _ := ret[0].(Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkCreateTx indicates an expected call of BulkCreateTx.
func (mr *MockRepositoryMockRecorder) BulkCreateTx(ctx, images interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateTx", reflect.TypeOf((*MockRepository)(nil).BulkCreateTx), ctx, images)
}

// Get mocks base method.
func (m *MockRepository) Get(ctx context.Context, filter *Filter) (*Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, filter)
	ret0, _ := ret[0].(*Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), ctx, filter)
}