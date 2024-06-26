package address

import (
	"context"
	"fmt"
)

//go:generate go run go.uber.org/mock/mockgen@v0.4.0 -source $GOFILE -destination "service_mock.go" -package "address"

// Service represents address use cases.
type Service interface {
	Get(ctx context.Context, filter *Filter) (*Address, error)
	Query(ctx context.Context, criteria *QueryCriteria) (Addresses, error)
	Create(ctx context.Context, props *Props) (*Address, error)
}

// UseCase represents address service.
type UseCase struct {
	addressRepo Repository
}

// NewUseCase create instance of UseCase.
func NewUseCase(addressRepo Repository) *UseCase {
	return &UseCase{addressRepo: addressRepo}
}

// Create address.
func (s UseCase) Create(ctx context.Context, props *Props) (*Address, error) {
	result, err := s.addressRepo.Create(ctx, props)
	if err != nil {
		return nil, fmt.Errorf("create address error: %w", err)
	}
	return result, nil
}

// Get address.
func (s UseCase) Get(ctx context.Context, filter *Filter) (*Address, error) {
	result, err := s.addressRepo.Get(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get address error: %w", err)
	}
	return result, nil
}

// Query addresses.
func (s UseCase) Query(ctx context.Context, criteria *QueryCriteria) (Addresses, error) {
	result, err := s.addressRepo.Query(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("query addresses error: %w", err)
	}
	return result, nil
}
