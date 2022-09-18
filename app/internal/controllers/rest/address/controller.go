package address

import (
	"context"
	"net/http"

	"github.com/blackPavlin/shop/app/internal/controllers/rest"
	"github.com/blackPavlin/shop/app/internal/controllers/rest/middlewares"
	"github.com/blackPavlin/shop/app/internal/entities"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddressUseCase
type AddressUseCase interface {
	Create(ctx context.Context, userID entities.UserID, address *entities.Address) (*entities.Address, error)
	FindByUserID(ctx context.Context, userID entities.UserID) (entities.Addresses, error)
}

// AddressController
type AddressController struct {
	addressUseCase AddressUseCase
	authMiddleware middlewares.AuthMiddleware
}

// NewAddressController
func NewAddressController(
	addressUseCase AddressUseCase,
	authMiddleware middlewares.AuthMiddleware,
) AddressController {
	return AddressController{
		addressUseCase: addressUseCase,
		authMiddleware: authMiddleware,
	}
}

// RegisterRoutes
func (ctrl AddressController) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/address")
	{
		group.Use(ctrl.authMiddleware.AuthGuard).GET("", ctrl.getAddresses)
		group.Use(ctrl.authMiddleware.AuthGuard).POST("", ctrl.createAddress)
		group.Use(ctrl.authMiddleware.AuthGuard).DELETE("", ctrl.deleteAddress)
	}
}

// getAddresses
func (ctrl AddressController) getAddresses(c *gin.Context) {
	userID, err := ctrl.authMiddleware.GetUserID(c)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	addresses, err := ctrl.addressUseCase.FindByUserID(c, userID)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, rest.DomainAddressesToResponse(addresses))
}

// createAddress
func (ctrl AddressController) createAddress(c *gin.Context) {
	req := &rest.CreateAddressRequest{}

	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := ctrl.authMiddleware.GetUserID(c)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	address := &entities.Address{
		UserID:   primitive.ObjectID(userID),
		City:     req.City,
		Country:  req.Country,
		Flat:     req.Flat,
		House:    req.House,
		Letter:   req.Letter,
		Postcode: req.Postcode,
		Street:   req.Street,
	}

	addr, err := ctrl.addressUseCase.Create(c, userID, address)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, rest.DomainAddressToResponse(addr))
}

// deleteAddress
func (ctrl AddressController) deleteAddress(c *gin.Context) {}
