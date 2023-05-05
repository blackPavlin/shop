// Code generated by MockGen. DO NOT EDIT.
// Source: product_repository.go

// Package order is a generated GoMock package.
package order

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// BulkCreateTx mocks base method.
func (m *MockProductRepository) BulkCreateTx(ctx context.Context) (Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateTx", ctx)
	ret0, _ := ret[0].(Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkCreateTx indicates an expected call of BulkCreateTx.
func (mr *MockProductRepositoryMockRecorder) BulkCreateTx(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateTx", reflect.TypeOf((*MockProductRepository)(nil).BulkCreateTx), ctx)
}
