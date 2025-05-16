package database

import (
	"context"
	"database/sql"
)

type TransactionManager interface {
	RunTransaction(ctx context.Context, options *sql.TxOptions, fn func(ctx context.Context) error) error
}

type NopTransactionManager struct{}

func NewNopTransactionManager() *NopTransactionManager {
	return &NopTransactionManager{}
}

func (NopTransactionManager) RunTransaction(ctx context.Context, _ *sql.TxOptions, fn func(context.Context) error) error {
	return fn(ctx)
}
