package product

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "product"

// Service
type Service interface {
}

// UseCase
type UseCase struct {
	productRepo Repository
}

// NewUseCase
func NewUseCase(productRepo Repository) *UseCase {
	return &UseCase{productRepo: productRepo}
}
