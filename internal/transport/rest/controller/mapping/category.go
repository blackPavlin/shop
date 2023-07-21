package mapping

import (
	"github.com/blackPavlin/shop/internal/domain/category"
	"github.com/blackPavlin/shop/internal/transport/rest"
)

// CreateCategoryResponse transform domain entity to rest response.
func CreateCategoryResponse(categ *category.Category) rest.Category {
	return rest.Category{
		Id:   int64(categ.ID),
		Name: categ.Props.Name,
	}
}

// CreateGetCategoriesResponse transform domain entity to rest response.
func CreateGetCategoriesResponse(categories category.Categories) rest.CategoryList {
	result := make(rest.CategoryList, 0, len(categories))
	for _, categ := range categories {
		result = append(result, CreateCategoryResponse(categ))
	}
	return result
}
