package usecases

import (
	"github.com/blackPavlin/shop/app/internal/core/usecases/auth"
	"github.com/blackPavlin/shop/app/internal/core/usecases/cart"
	"github.com/blackPavlin/shop/app/internal/core/usecases/user"
)

type UseCases struct {
	AuthUseCase auth.AuthUseCase
	UserUseCase user.UserUseCase
	CartUseCase cart.CartUseCase
}
