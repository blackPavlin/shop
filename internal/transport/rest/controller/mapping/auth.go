package mapping

import "github.com/blackPavlin/shop/internal/transport/rest"

// CreateLoginResponse transform domain entity to rest response.
func CreateLoginResponse(token string) rest.LoginResponse {
	return rest.LoginResponse{
		Token: token,
	}
}
