package controller

import (
	"io"
	"net/http"

	"github.com/blackPavlin/shop/internal/domain/image"
	"github.com/blackPavlin/shop/internal/transport/rest/controller/mapping"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/restx"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// ImageController
type ImageController struct {
	imageService image.Service
}

// NewImageController
func NewImageController(imageService image.Service) *ImageController {
	return &ImageController{imageService: imageService}
}

// RegisterRoutes
func (ctrl *ImageController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/image", func(r chi.Router) {
		r.Post("/upload", ctrl.UploadImageHandler)
	})
}

// UploadImageHandler
func (ctrl *ImageController) UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(5 << 20); err != nil {
		restx.HandleError(w, r, err)
		return
	}
	files, ok := r.MultipartForm.File["files"]
	if !ok || len(files) == 0 {
		restx.HandleError(w, r, errorx.NewBadRequestError("files required"))
		return
	}
	images := make([]*image.StorageProps, 0, len(files))
	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			restx.HandleError(w, r, errorx.ErrInternal)
			return
		}
		defer f.Close()
		buff, err := io.ReadAll(f)
		if err != nil {
			restx.HandleError(w, r, errorx.ErrInternal)
			return
		}
		images = append(images, &image.StorageProps{
			Name:   file.Filename,
			Buffer: buff,
		})
	}
	result, err := ctrl.imageService.BulkCreate(r.Context(), images)
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Respond(w, r, mapping.CreateUploadImagesResponse(result))
}
