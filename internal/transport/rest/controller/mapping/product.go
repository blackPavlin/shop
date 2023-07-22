package mapping

import (
	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/internal/transport/rest"
)

// CreateProductResponse transform domain entity to rest response.
func CreateProductResponse(p *product.Product) rest.Product {
	return rest.Product{
		Id:          int64(p.ID),
		CategoryId:  int64(p.Props.CategoryID),
		Name:        p.Props.Name,
		Description: &p.Props.Description,
		Amount:      p.Props.Amount,
		Price:       p.Props.Price,
	}
}

// CreateGetProductsResponse transform domain entity to rest response.
func CreateGetProductsResponse(products product.Products) {}
