package mapping

import (
	"github.com/blackPavlin/shop/internal/domain/image"
	"github.com/blackPavlin/shop/internal/transport/rest"
)

// CreateUploadImageResponse
func CreateUploadImageResponse(image *image.Image) rest.Image {
	return rest.Image{
		Id:           int64(image.ID),
		Name:         image.Props.Name,
		OriginalName: image.Props.OriginalName,
	}
}

// CreateUploadImagesResponse
func CreateUploadImagesResponse(images image.Images) rest.ImageList {
	result := make(rest.ImageList, 0, len(images))
	for _, image := range images {
		result = append(result, CreateUploadImageResponse(image))
	}
	return result
}
