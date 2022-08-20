package transaction

import "context"

// TransactionService
type TransactionService struct {
	transactionRepository TransactionRepository
}

// NewTransactionService
func NewTransactionService(transactionRepository TransactionRepository) TransactionService {
	return TransactionService{
		transactionRepository: transactionRepository,
	}
}

// RunTransaction
func (s TransactionService) RunTransaction(ctx context.Context, callback func(context context.Context) error) error {
	return s.transactionRepository.RunTransaction(ctx, callback)
}
