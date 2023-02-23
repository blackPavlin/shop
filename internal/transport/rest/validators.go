package rest

import (
	"github.com/blackPavlin/shop/pkg/errorx"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Validate
func (r LoginRequest) Validate() error {
	if err := validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.EmailFormat),
		validation.Field(&r.Password, validation.Required),
	); err != nil {
		return errorx.NewBadRequestError(err.Error())
	}
	return nil
}

// Validate
func (r SignupRequest) Validate() error {
	if err := validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.EmailFormat),
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Phone, validation.Required, is.E164),
		validation.Field(&r.Password, validation.Required),
	); err != nil {
		return errorx.NewBadRequestError(err.Error())
	}
	return nil
}

// Validate
func (r CreateAddressRequest) Validate() error {
	if err := validation.ValidateStruct(&r,
		validation.Field(&r.City, validation.Required),
		validation.Field(&r.Country, validation.Required),
		validation.Field(&r.House, validation.Required, validation.Min(1)),
		validation.Field(&r.Postcode, validation.Required, validation.Min(1)),
		validation.Field(&r.Street, validation.Required),
	); err != nil {
		return errorx.NewBadRequestError(err.Error())
	}
	return nil
}

// Validate
func (r Category) Validate() error {
	if err := validation.ValidateStruct(&r,
		validation.Field(&r.Id, validation.Required),
		validation.Field(&r.Name, validation.Required),
	); err != nil {
		return errorx.NewBadRequestError(err.Error())
	}
	return nil
}

// Validate
func (r CreateCategoryRequest) Validate() error {
	if err := validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required),
	); err != nil {
		return errorx.NewBadRequestError(err.Error())
	}
	return nil
}
