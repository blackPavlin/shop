package mapping

import (
	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/internal/transport/rest"
)

// CreateGetUserResponse transform domain entity to rest response.
func CreateGetUserResponse(u *user.User) rest.User {
	return rest.User{
		Id:    int64(u.ID),
		Role:  rest.UserRole(u.Role.String()),
		Email: u.Props.Email,
		Name:  u.Props.Name,
		Phone: u.Props.Phone,
	}
}
