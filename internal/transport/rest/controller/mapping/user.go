package mapping

import (
	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/internal/transport/rest"
)

// CreateGetUserResponse
func CreateGetUserResponse(user *user.User) rest.User {
	return rest.User{
		Id:    int64(user.ID),
		Role:  rest.UserRole(user.Role.String()),
		Email: user.Props.Email,
		Name:  user.Props.Name,
		Phone: user.Props.Phone,
	}
}
