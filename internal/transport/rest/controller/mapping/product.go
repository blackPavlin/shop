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
		Images:      CreateUploadImagesResponse(p.Images),
	}
}

// CreateProductsResponse transform domain entity to rest response.
func CreateProductsResponse(products product.Products) []rest.Product {
	result := make([]rest.Product, 0, len(products))
	for _, p := range products {
		result = append(result, CreateProductResponse(p))
	}
	return result
}

// CreateGetProductsResponse transform domain entity to rest response.
func CreateGetProductsResponse(products *product.QueryResult) rest.ProductList {
	return rest.ProductList{
		Quantity: products.Count,
		Items:    CreateProductsResponse(products.Items),
	}
}
