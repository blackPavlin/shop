package product

// ProductService
type ProductService interface{}

// ProductUseCase
type ProductUseCase struct {
	productService ProductService
}

// NewProductUseCase
func NewProductUseCase(productService ProductService) ProductUseCase {
	return ProductUseCase{
		productService: productService,
	}
}
