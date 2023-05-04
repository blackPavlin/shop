package pg

import (
	"context"
	"database/sql"
	"fmt"

	"go.uber.org/zap"

	"github.com/blackPavlin/shop/ent"
)

// TxManager
type TxManager struct {
	client *ent.Client
	logger *zap.Logger
}

// NewTxManager
func NewTxManager(client *ent.Client, logger *zap.Logger) *TxManager {
	return &TxManager{client: client, logger: logger}
}

// RunTransaction
func (p *TxManager) RunTransaction(
	ctx context.Context,
	options *sql.TxOptions,
	fn func(ctx context.Context) error,
) error {
	tx, err := p.client.BeginTx(ctx, options)
	if err != nil {
		return fmt.Errorf("start transaction error: %w", err)
	}
	ctx = ent.NewTxContext(ctx, tx)
	defer func() {
		if v := recover(); v != nil {
			err = tx.Rollback()
			if err != nil {
				p.logger.Error(
					"transaction rollback error",
					zap.Any("panic", v),
					zap.NamedError("txErr", err),
				)
			}
			panic(v)
		}
	}()
	if err := fn(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rolling back transaction error: %w", err)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction error: %w", err)
	}
	return nil
}
