package mapping

import "github.com/blackPavlin/shop/internal/transport/rest"

// CreateLoginResponse
func CreateLoginResponse(token string) rest.LoginResponse {
	return rest.LoginResponse{
		Token: token,
	}
}
