package controller

import (
	"net/http"

	"github.com/blackPavlin/shop/internal/domain/order"
	"github.com/blackPavlin/shop/internal/transport/rest/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// OrderController represents order controller.
type OrderController struct {
	orderService   order.Service
	authMiddleware *middleware.Middleware
}

// NewOrderController create instance of OrderController.
func NewOrderController(
	orderService order.Service,
	authMiddleware *middleware.Middleware,
) *OrderController {
	return &OrderController{orderService, authMiddleware}
}

// RegisterRoutes register routes to the specified router.
func (ctrl *OrderController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/order", func(r chi.Router) {
		r.Use(ctrl.authMiddleware.Authorization)
		r.Get("/", ctrl.GetOrdersHandler)
		r.Post("/", ctrl.CreateOrderHandler)
		r.Route("/{orderID}", func(r chi.Router) {
			r.Get("/", ctrl.GetOrderHandler)
		})
	})
}

// GetOrdersHandler define handler for GET /api/order.
func (ctrl *OrderController) GetOrdersHandler(w http.ResponseWriter, r *http.Request) {}

// CreateOrderHandler define handler for POST /api/order.
func (ctrl *OrderController) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusCreated)
}

// GetOrderHandler define handler for GET /api/order/{orderId}.
func (ctrl *OrderController) GetOrderHandler(w http.ResponseWriter, r *http.Request) {}
