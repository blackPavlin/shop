package controller

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/blackPavlin/shop/internal/domain/image"
	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/internal/transport/rest"
	"github.com/blackPavlin/shop/internal/transport/rest/controller/mapping"
	"github.com/blackPavlin/shop/internal/transport/rest/middleware"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/restx"
)

const maxFileSize = 5 << 20

// ProductController represents product controller.
type ProductController struct {
	productService product.Service
	imageService   product.ImageService
	authMiddleware *middleware.Middleware
}

// NewProductController create instance of ProductController.
func NewProductController(
	productService product.Service,
	imageService product.ImageService,
	authMiddleware *middleware.Middleware,
) *ProductController {
	return &ProductController{productService, imageService, authMiddleware}
}

// RegisterRoutes register routes to the specified router.
func (ctrl *ProductController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/product", func(r chi.Router) {
		r.Get("/", ctrl.GetProductsHandler)
		r.Group(func(r chi.Router) {
			r.Use(ctrl.authMiddleware.Authorization)
			r.Post("/", ctrl.CreateProductHandler)
		})
		r.Route("/{productID}", func(r chi.Router) {
			r.Get("/", ctrl.GetProductHandler)
			r.Group(func(r chi.Router) {
				r.Use(ctrl.authMiddleware.Authorization)
				r.Patch("/", ctrl.UpdateProductHandler)
			})
			r.Route("/image", func(r chi.Router) {
				r.Use(ctrl.authMiddleware.Authorization)
				r.Post("/", ctrl.UploadProductImageHandler)
				r.Delete("/{imageID}", ctrl.RemoveProductImageHandler)
			})
		})
	})
}

// GetProductsHandler define handler for GET /api/product.
func (ctrl *ProductController) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := ctrl.productService.Query(r.Context(), &product.QueryCriteria{})
	if err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	render.Respond(w, r, mapping.CreateGetProductsResponse(products))
}

// CreateProductHandler define handler for POST /api/product.
func (ctrl *ProductController) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	request := &rest.CreateProductRequest{}
	if err := render.DecodeJSON(r.Body, request); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	if err := request.Validate(); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	p, err := ctrl.productService.Create(r.Context(), request.ToDomainEntity())
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Status(r, http.StatusCreated)
	render.Respond(w, r, mapping.CreateProductResponse(p))
}

// GetProductHandler define handler for GET /api/product/{productId}.
func (ctrl *ProductController) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	productID, err := restx.GetIDFromURLParams(r, "productID")
	if err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	p, err := ctrl.productService.Get(r.Context(), &product.Filter{
		ID: product.IDFilter{Eq: product.IDs{product.ID(productID)}},
	})
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Respond(w, r, mapping.CreateProductResponse(p))
}

// UpdateProductHandler define handler for PATCH /api/product/{productId}.
func (ctrl *ProductController) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	productID, err := restx.GetIDFromURLParams(r, "productID")
	if err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	request := &rest.CreateProductRequest{}
	if err := render.DecodeJSON(r.Body, request); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	if err := request.Validate(); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	p, err := ctrl.productService.Update(
		r.Context(),
		product.ID(productID),
		request.ToDomainEntity(),
	)
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Respond(w, r, mapping.CreateProductResponse(p))
}

// UploadProductImageHandler define handler for POST /api/product/{productId}/image.
func (ctrl *ProductController) UploadProductImageHandler(w http.ResponseWriter, r *http.Request) {
	productID, err := restx.GetIDFromURLParams(r, "productID")
	if err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	if err := r.ParseMultipartForm(maxFileSize); err != nil {
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
	result, err := ctrl.imageService.BulkCreate(r.Context(), product.ID(productID), images)
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Status(r, http.StatusCreated)
	render.Respond(w, r, mapping.CreateUploadImagesResponse(result))
}

// RemoveProductImageHandler define handler for DELETE /api/product/{productId}/image/{imageID}.
func (ctrl *ProductController) RemoveProductImageHandler(w http.ResponseWriter, r *http.Request) {
	imageID, err := restx.GetIDFromURLParams(r, "imageID")
	if err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	if err := ctrl.imageService.Delete(r.Context(), product.ImageID(imageID)); err != nil {
		restx.HandleError(w, r, err)
		return
	}
}
