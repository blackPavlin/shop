package usecases

import (
	"github.com/blackPavlin/shop/app/internal/usecases/auth"
	"github.com/blackPavlin/shop/app/internal/usecases/cart"
	"github.com/blackPavlin/shop/app/internal/usecases/category"
	"github.com/blackPavlin/shop/app/internal/usecases/user"
)

type UseCases struct {
	AuthUseCase     auth.AuthUseCase
	UserUseCase     user.UserUseCase
	CartUseCase     cart.CartUseCase
	CategoryUseCase category.CategoryUseCase
}
