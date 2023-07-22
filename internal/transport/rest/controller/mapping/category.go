package mapping

import (
	"github.com/blackPavlin/shop/internal/domain/category"
	"github.com/blackPavlin/shop/internal/transport/rest"
)

// CreateCategoryResponse transform domain entity to rest response.
func CreateCategoryResponse(c *category.Category) rest.Category {
	return rest.Category{
		Id:   int64(c.ID),
		Name: c.Props.Name,
	}
}

// CreateGetCategoriesResponse transform domain entity to rest response.
func CreateGetCategoriesResponse(categories category.Categories) rest.CategoryList {
	result := make(rest.CategoryList, 0, len(categories))
	for _, c := range categories {
		result = append(result, CreateCategoryResponse(c))
	}
	return result
}
