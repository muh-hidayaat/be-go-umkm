package request

import (
	"github.com/go-playground/validator/v10"
)

// Initialize the validator instance
var validate = validator.New()

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UpdatePasswordRequest struct {
	PasswordOld     string `form:"password_old" validate:"required"`
	Password        string `form:"password" validate:"required,min=8"`
	ConfirmPassword string `form:"confirm_password" validate:"required,eqfield=Password"`
}

func (r *CreateUserRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateUserRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdatePasswordRequest) Validate() error {
	return validate.Struct(r)
}
