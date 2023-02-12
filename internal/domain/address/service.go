package address

import "context"

// AddressService
type AddressService interface {
	Create(ctx context.Context, props *Props) (*Address, error)
	Get(ctx context.Context, filter *Filter) (*Address, error)
	Query(ctx context.Context, criteria *QueryCriteria) (Addresses, error)
}

// AddressUseCase
type AddressUseCase struct {
	addressRepo Repository
}

// NewAddressUseCase
func NewAddressUseCase(addressRepo Repository) *AddressUseCase {
	return &AddressUseCase{addressRepo: addressRepo}
}

// Create
func (s *AddressUseCase) Create(ctx context.Context, props *Props) (*Address, error) {
	return s.addressRepo.Create(ctx, props)
}

// Get
func (s *AddressUseCase) Get(ctx context.Context, filter *Filter) (*Address, error) {
	return s.addressRepo.Get(ctx, filter)
}

// Query
func (s *AddressUseCase) Query(ctx context.Context, criteria *QueryCriteria) (Addresses, error) {
	return s.addressRepo.Query(ctx, criteria)
}
