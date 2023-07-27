package rest

import (
	"github.com/blackPavlin/shop/internal/domain/address"
	"github.com/blackPavlin/shop/internal/domain/auth"
	"github.com/blackPavlin/shop/internal/domain/category"
	"github.com/blackPavlin/shop/internal/domain/product"
)

// ToDomainEntity transforms api template into domain entity.
func (r LoginRequest) ToDomainEntity() *auth.LoginProps {
	return &auth.LoginProps{
		Email:    r.Email,
		Password: r.Password,
	}
}

// ToDomainEntity transforms api template into domain entity.
func (r SignupRequest) ToDomainEntity() *auth.SignupProps {
	return &auth.SignupProps{
		Email:    r.Email,
		Name:     r.Name,
		Phone:    r.Phone,
		Password: r.Password,
	}
}

// ToDomainEntity transforms api template into domain entity.
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

// ToDomainEntity transforms api template into domain entity.
func (r CreateCategoryRequest) ToDomainEntity() *category.Props {
	return &category.Props{
		Name: r.Name,
	}
}

// ToDomainEntity transforms api template into domain entity.
func (r CreateProductRequest) ToDomainEntity() *product.Props {
	return &product.Props{
		CategoryID:  category.ID(r.CategoryId),
		Name:        r.Name,
		Description: r.Description,
		Amount:      r.Amount,
		Price:       r.Price,
	}
}
