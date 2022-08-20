// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package transaction is a generated GoMock package.
package transaction

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// RunTransaction mocks base method.
func (m *MockTransactionRepository) RunTransaction(ctx context.Context, callback func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTransaction", ctx, callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTransaction indicates an expected call of RunTransaction.
func (mr *MockTransactionRepositoryMockRecorder) RunTransaction(ctx, callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTransaction", reflect.TypeOf((*MockTransactionRepository)(nil).RunTransaction), ctx, callback)
}