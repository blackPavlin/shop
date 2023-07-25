package rest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/blackPavlin/shop/pkg/errorx"
)

// Validate body of LoginRequest.
func (r LoginRequest) Validate() error {
	if err := validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.EmailFormat),
		validation.Field(&r.Password, validation.Required),
	); err != nil {
		return errorx.NewBadRequestError(err.Error())
	}
	return nil
}

// Validate body of SignupRequest.
func (r SignupRequest) Validate() error {
	if err := validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.EmailFormat),
		validation.Field(&r.Name, validation.Required, validation.Length(1, 100)),
		validation.Field(&r.Phone, validation.Required, is.E164, validation.Length(1, 100)),
		validation.Field(&r.Password, validation.Required, validation.Length(1, 100)),
	); err != nil {
		return errorx.NewBadRequestError(err.Error())
	}
	return nil
}

// Validate body of CreateAddressRequest.
func (r CreateAddressRequest) Validate() error {
	if err := validation.ValidateStruct(&r,
		validation.Field(&r.City, validation.Required),
		validation.Field(&r.Country, validation.Required),
		validation.Field(&r.House, validation.Required, validation.Min(1)),
		validation.Field(&r.Postcode, validation.Required, validation.Min(1)),
		validation.Field(&r.Street, validation.Required, validation.Length(1, 50)),
	); err != nil {
		return errorx.NewBadRequestError(err.Error())
	}
	return nil
}

// Validate body of Category.
func (r Category) Validate() error {
	if err := validation.ValidateStruct(&r,
		validation.Field(&r.Id, validation.Required),
		validation.Field(&r.Name, validation.Required, validation.Length(1, 50)),
	); err != nil {
		return errorx.NewBadRequestError(err.Error())
	}
	return nil
}

// Validate body of CreateCategoryRequest.
func (r CreateCategoryRequest) Validate() error {
	if err := validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.Length(1, 50)),
	); err != nil {
		return errorx.NewBadRequestError(err.Error())
	}
	return nil
}

// Validate body of CreateProductRequest.
func (r CreateProductRequest) Validate() error {
	if err := validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Description, validation.Required),
		validation.Field(&r.CategoryId, validation.Required),
		validation.Field(&r.Price, validation.Required, validation.Min(0)),
		validation.Field(&r.Amount, validation.Required, validation.Min(0)),
	); err != nil {
		return errorx.NewBadRequestError(err.Error())
	}
	return nil
}
