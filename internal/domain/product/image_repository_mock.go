// Code generated by MockGen. DO NOT EDIT.
// Source: image_repository.go
//
// Generated by this command:
//
//	mockgen -source image_repository.go -destination image_repository_mock.go -package product
//

// Package product is a generated GoMock package.
package product

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
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

// BulkCreate mocks base method.
func (m *MockImageRepository) BulkCreate(ctx context.Context, images Images) (Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreate", ctx, images)
	ret0, _ := ret[0].(Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkCreate indicates an expected call of BulkCreate.
func (mr *MockImageRepositoryMockRecorder) BulkCreate(ctx, images any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreate", reflect.TypeOf((*MockImageRepository)(nil).BulkCreate), ctx, images)
}

// Delete mocks base method.
func (m *MockImageRepository) Delete(ctx context.Context, filter *ImageFilter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockImageRepositoryMockRecorder) Delete(ctx, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockImageRepository)(nil).Delete), ctx, filter)
}

// Get mocks base method.
func (m *MockImageRepository) Get(ctx context.Context, filter *ImageFilter) (*Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, filter)
	ret0, _ := ret[0].(*Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockImageRepositoryMockRecorder) Get(ctx, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockImageRepository)(nil).Get), ctx, filter)
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
func (mr *MockImageRepositoryMockRecorder) Query(ctx, criteria any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockImageRepository)(nil).Query), ctx, criteria)
}
