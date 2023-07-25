// Code generated by MockGen. DO NOT EDIT.
// Source: image_repository.go

// Package product is a generated GoMock package.
package product

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockImageRepository is a mock of ImageRepository interface.
type MockImageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockImageRepositoryMockRecorder
}

// MockImageRepositoryMockRecorder is the mock recorder for MockImageRepository.
type MockImageRepositoryMockRecorder struct {
	mock *MockImageRepository
}

// NewMockImageRepository creates a new mock instance.
func NewMockImageRepository(ctrl *gomock.Controller) *MockImageRepository {
	mock := &MockImageRepository{ctrl: ctrl}
	mock.recorder = &MockImageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageRepository) EXPECT() *MockImageRepositoryMockRecorder {
	return m.recorder
}

// BulkCreateTx mocks base method.
func (m *MockImageRepository) BulkCreateTx(ctx context.Context, images Images) (Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateTx", ctx, images)
	ret0, _ := ret[0].(Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkCreateTx indicates an expected call of BulkCreateTx.
func (mr *MockImageRepositoryMockRecorder) BulkCreateTx(ctx, images interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateTx", reflect.TypeOf((*MockImageRepository)(nil).BulkCreateTx), ctx, images)
}

// Query mocks base method.
func (m *MockImageRepository) Query(ctx context.Context, criteria *ImageQueryCriteria) (Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", ctx, criteria)
	ret0, _ := ret[0].(Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockImageRepositoryMockRecorder) Query(ctx, criteria interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockImageRepository)(nil).Query), ctx, criteria)
}
