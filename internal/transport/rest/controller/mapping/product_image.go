package mapping

import (
	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/internal/transport/rest"
)

// CreateUploadImageResponse transform domain entity to rest response.
func CreateUploadImageResponse(img *product.Image) rest.Image {
	return rest.Image{
		Id:           int64(img.ID),
		// todo: product_id
		Name:         img.Props.Name,
	}
}

// CreateUploadImagesResponse transform domain entity to rest response.
func CreateUploadImagesResponse(images product.Images) rest.ImageList {
	result := make(rest.ImageList, 0, len(images))
	for _, img := range images {
		result = append(result, CreateUploadImageResponse(img))
	}
	return result
}
