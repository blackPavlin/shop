package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// TransactionRepository
type TransactionRepository struct {
	db *mongo.Database
}

// NewTransactionRepository
func NewTransactionRepository(db *mongo.Database) TransactionRepository {
	return TransactionRepository{
		db: db,
	}
}

// RunTransaction
func (t TransactionRepository) RunTransaction(
	ctx context.Context,
	callback func(context context.Context) error,
) error {
	return t.db.Client().UseSession(ctx, func(sctx mongo.SessionContext) error {
		if err := sctx.StartTransaction(); err != nil {
			return err
		}

		if err := callback(sctx); err != nil {
			if err := sctx.AbortTransaction(sctx); err != nil {
				return err
			}

			return err
		}

		return sctx.CommitTransaction(sctx)
	})
}
