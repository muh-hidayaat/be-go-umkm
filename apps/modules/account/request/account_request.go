package request

import "github.com/go-playground/validator/v10"

type AccountCreateRequest struct {
	Name    string `json:"name" validate:"required" example:"Account A"`
	Type    string `json:"type" validate:"required" example:"E-Wallet"`
	Balance int    `json:"balance" validate:"required" example:"110000"`
}

type AccountUpdateRequest struct {
	Name    string `json:"name" validate:"omitempty" example:"Account A"`
	Type    string `json:"type" validate:"omitempty" example:"E-Wallet"`
	Balance int    `json:"balance" validate:"omitempty" example:"110000"`
}

var validateAccount = validator.New()

func (r *AccountCreateRequest) Validate() error {
	return validateAccount.Struct(r)
}

func (r *AccountUpdateRequest) Validate() error {
	return validateAccount.Struct(r)
}
