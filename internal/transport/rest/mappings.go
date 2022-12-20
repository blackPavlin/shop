package rest

import "github.com/blackPavlin/shop/internal/domain/auth"

// ToDomainEntity
func (r LoginRequest) ToDomainEntity() *auth.LoginProps {
	return &auth.LoginProps{
		Email:    r.Email,
		Password: r.Password,
	}
}

// ToDomainEntity
func (r SignupRequest) ToDomainEntity() *auth.SignupProps {
	return &auth.SignupProps{
		Email:    r.Email,
		Name:     r.Name,
		Phone:    r.Phone,
		Password: r.Password,
	}
}
