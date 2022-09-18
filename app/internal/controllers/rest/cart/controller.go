package cart

import (
	"context"
	"net/http"

	"github.com/blackPavlin/shop/app/internal/controllers/rest"
	"github.com/blackPavlin/shop/app/internal/controllers/rest/middlewares"
	"github.com/blackPavlin/shop/app/internal/entities"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CartUseCase
type CartUseCase interface {
	GetCartByUserID(ctx context.Context, userID entities.UserID) (*entities.Cart, error)
	ClearCartByUserID(ctx context.Context, userID entities.UserID) (*entities.Cart, error)
	AppendProduct(ctx context.Context, userID entities.UserID, product *entities.CartProduct) (*entities.Cart, error)
	UpdateProduct(ctx context.Context, userID entities.UserID, product *entities.CartProduct) (*entities.Cart, error)
}

// CartController
type CartController struct {
	cartUseCase    CartUseCase
	authMiddleware middlewares.AuthMiddleware
}

// NewCartController
func NewCartController(cartUseCase CartUseCase, authMiddleware middlewares.AuthMiddleware) CartController {
	return CartController{
		cartUseCase:    cartUseCase,
		authMiddleware: authMiddleware,
	}
}

// RegisterRoutes
func (ctrl CartController) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/cart")
	{
		group.Use(ctrl.authMiddleware.AuthGuard).GET("", ctrl.getCart)
		group.Use(ctrl.authMiddleware.AuthGuard).POST("", ctrl.appendProduct)
		group.Use(ctrl.authMiddleware.AuthGuard).PATCH("", ctrl.updateProduct)
		group.Use(ctrl.authMiddleware.AuthGuard).DELETE("", ctrl.clearCart)
	}
}

// getCart
func (ctrl CartController) getCart(c *gin.Context) {
	userID, err := ctrl.authMiddleware.GetUserID(c)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	cart, err := ctrl.cartUseCase.GetCartByUserID(c, userID)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, rest.DomainCartToResponse(cart))
}

// appendProduct
func (ctrl CartController) appendProduct(c *gin.Context) {
	req := &rest.CartProduct{}

	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := ctrl.authMiddleware.GetUserID(c)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	productID, err := primitive.ObjectIDFromHex(req.ProductId)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	product := &entities.CartProduct{
		ProductID: productID,
		Amount:    int64(req.Amount),
	}

	cart, err := ctrl.cartUseCase.AppendProduct(c, userID, product)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, rest.DomainCartToResponse(cart))
}

// updateProduct
func (ctrl CartController) updateProduct(c *gin.Context) {
	req := &rest.CartProduct{}

	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := ctrl.authMiddleware.GetUserID(c)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	productID, err := primitive.ObjectIDFromHex(req.ProductId)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	product := &entities.CartProduct{
		ProductID: productID,
		Amount:    int64(req.Amount),
	}

	cart, err := ctrl.cartUseCase.UpdateProduct(c, userID, product)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, rest.DomainCartToResponse(cart))
}

// clearCart
func (ctrl CartController) clearCart(c *gin.Context) {
	userID, err := ctrl.authMiddleware.GetUserID(c)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	cart, err := ctrl.cartUseCase.ClearCartByUserID(c, userID)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, rest.DomainCartToResponse(cart))
}
