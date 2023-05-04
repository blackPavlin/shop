package repositoryx

import (
	"context"
	"database/sql"
)

// TxManager is repository transaction manager. If repository implementation not support transaction use NopTxManager.
type TxManager interface {
	RunTransaction(ctx context.Context, options *sql.TxOptions, fn func(ctx context.Context) error) error
}

// NopTxManager transaction manager for tests.
type NopTxManager struct{}

// RunTransaction is no-op.
func (NopTxManager) RunTransaction(ctx context.Context, _ *sql.TxOptions, fn func(context.Context) error) error {
	return fn(ctx)
}

// NewNopTxManager create new NopTxManager.
func NewNopTxManager() *NopTxManager { return &NopTxManager{} }
