package address

import "context"

// Service
type Service interface {
	Create(ctx context.Context, props *Props) (*Address, error)
	Get(ctx context.Context, filter *Filter) (*Address, error)
	Query(ctx context.Context, criteria *QueryCriteria) (Addresses, error)
}

// UseCase
type UseCase struct {
	addressRepo Repository
}

// NewUseCase
func NewUseCase(addressRepo Repository) *UseCase {
	return &UseCase{addressRepo: addressRepo}
}

// Create
func (s *UseCase) Create(ctx context.Context, props *Props) (*Address, error) {
	return s.addressRepo.Create(ctx, props)
}

// Get
func (s *UseCase) Get(ctx context.Context, filter *Filter) (*Address, error) {
	return s.addressRepo.Get(ctx, filter)
}

// Query
func (s *UseCase) Query(ctx context.Context, criteria *QueryCriteria) (Addresses, error) {
	return s.addressRepo.Query(ctx, criteria)
}
