package mapping

import (
	"github.com/blackPavlin/shop/internal/domain/category"
	"github.com/blackPavlin/shop/internal/transport/rest"
)

// CreateGetCategoriesResponse
func CreateGetCategoriesResponse(categories category.Categories) rest.CategoryList {
	result := make(rest.CategoryList, 0, len(categories))

	return result
}
