package helpers

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteJSON(ctx *fiber.Ctx, statusCode int, data interface{}, message string) error {
	// Set the status code and respond with JSON
	return ctx.Status(statusCode).JSON(Response{
		Status:  statusCode,
		Message: message,
		Data:    data,
	})
}
