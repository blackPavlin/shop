package mapping

import (
	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/internal/transport/rest"
)

// CreateGetUserResponse
func CreateGetUserResponse(user *user.User) rest.User {
	return rest.User{}
}
