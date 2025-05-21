package request

import (
	"github.com/go-playground/validator/v10"
)

// Initialize the validator instance
var validate = validator.New()

// UserLoginRequest defines the structure and validation rules for user login
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UpdatePasswordRequest struct {
	PasswordOld     string `json:"password_old" validate:"required" example:"oldpassword123"`
	Password        string `json:"password" validate:"required,min=8" example:"newpassword123"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8" example:"newpassword123"`
}

type UpdateEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type UpdatePhoneRequest struct {
	Phone string `json:"phone" validate:"required,min=8"`
}

type UpdateNameRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
}

type UpdateFavRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type UpdatePhotoRequest struct {
	Photo string `json:"photo" validate:"required,endswith=jpg|endswith=jpeg|endswith=png,maxsize=3145728"`
}

// **CreateUserRequest**
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// **UpdateUserRequest**
type UpdateUserRequest struct {
	Name   string `json:"name" validate:"omitempty,min=2,max=100"`
	Email  string `json:"email" validate:"omitempty,email"`
	RoleID string `json:"role_id" validate:"required,uuid"`
}

// **CreateUserRequest**
type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100" example:"John Doe"`
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required,min=8" example:"password123"`
}

// Validate methods for each struct to check the validity of fields
func (r *UpdatePhotoRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateNameRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdatePhoneRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdatePasswordRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateEmailRequest) Validate() error {
	return validate.Struct(r)
}

func (r *CreateUserRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateUserRequest) Validate() error {
	return validate.Struct(r)
}
func (r *RegisterUserRequest) Validate() error {
	return validate.Struct(r)
}

// VerifyOTPRequest defines the structure and validation rules for OTP verification
type VerifyOTPRequest struct {
	Email   string `json:"email" validate:"required,email"`
	OTPCode string `json:"otp_code" validate:"required"`
}

// Validate checks the validity of the UserLoginRequest fields
func (r *UserLoginRequest) Validate() error {
	return validate.Struct(r)
}

// Validate checks the validity of the UserLoginRequest fields
func (r *UpdateFavRequest) Validate() error {
	return validate.Struct(r)
}

// Validate checks the validity of the VerifyOTPRequest fields
func (r *VerifyOTPRequest) Validate() error {
	return validate.Struct(r)
}

type SetNewPasswordRequest struct {
	Email    string `json:"email" validate:"required,email"`
	OTPCode  string `json:"otp_code" validate:"required,len=6"`
	Password string `json:"password" validate:"required,min=8"`
}

func (r *SetNewPasswordRequest) Validate() error {
	return validate.Struct(r)
}
