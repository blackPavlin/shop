package mapping

import (
	"github.com/blackPavlin/shop/internal/domain/image"
	"github.com/blackPavlin/shop/internal/transport/rest"
)

// CreateUploadImageResponse transform domain entity to rest response.
func CreateUploadImageResponse(img *image.Image) rest.Image {
	return rest.Image{
		Id:           int64(img.ID),
		Name:         img.Props.Name,
		OriginalName: img.Props.OriginalName,
	}
}

// CreateUploadImagesResponse transform domain entity to rest response.
func CreateUploadImagesResponse(images image.Images) rest.ImageList {
	result := make(rest.ImageList, 0, len(images))
	for _, img := range images {
		result = append(result, CreateUploadImageResponse(img))
	}
	return result
}
