package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	FindByID(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error

	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	// ForgotPassword(ctx *fiber.Ctx) error
	// VerifyEmail(ctx *fiber.Ctx) error
}
