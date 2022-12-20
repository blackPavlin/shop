package product

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "product"

// ProductService
type ProductService interface {
}

// ProductUseCase
type ProductUseCase struct {
	productRepo Repository
}

// NewProductUseCase
func NewProductUseCase(productRepo Repository) *ProductUseCase {
	return &ProductUseCase{productRepo: productRepo}
}
