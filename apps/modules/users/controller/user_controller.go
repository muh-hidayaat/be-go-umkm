package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	ChangePassword(ctx *fiber.Ctx) error
}
