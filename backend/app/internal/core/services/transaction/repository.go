package transaction

import "context"

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "transaction"

// TransactionRepository
type TransactionRepository interface {
	RunTransaction(ctx context.Context, callback func(context context.Context) error) error
}
