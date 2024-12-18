package validator

import (
	"beli-tanah/model/web"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type IUserValidator interface {
	ValidateRegisterUser(req web.RegisterUserRequest) error
	ValidateLoginUser(req web.LoginUserRequest) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) ValidateRegisterUser(req web.RegisterUserRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Name,
			validation.Required.Error("full name is required"),
		),
		validation.Field(
			&req.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("invalid email format"),
		),
		validation.Field(
			&req.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(8, 30).Error("password must be between 8 and 30 characters"),
		),
	)
}

func (uv *userValidator) ValidateLoginUser(req web.LoginUserRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("invalid email format"),
		),
		validation.Field(
			&req.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(8, 30).Error("password must be between 8 and 30 characters"),
		),
	)
}
