package rest

import (
	"github.com/blackPavlin/shop/internal/domain/address"
	"github.com/blackPavlin/shop/internal/domain/auth"
	"github.com/blackPavlin/shop/internal/domain/category"
)

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

// ToDomainEntity
func (r CreateAddressRequest) ToDomainEntity() *address.Props {
	return &address.Props{
		City:     r.City,
		Country:  r.Country,
		Flat:     r.Flat,
		House:    r.House,
		Letter:   r.Letter,
		Postcode: r.Postcode,
		Street:   r.Street,
	}
}

// ToDomainEntity
func (r CreateCategoryRequest) ToDomainEntity() *category.Props {
	return &category.Props{
		Name: r.Name,
	}
}