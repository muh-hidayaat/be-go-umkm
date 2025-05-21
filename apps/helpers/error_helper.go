package helpers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// HTTPError represents a structured error response for APIs.
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HTTPSuccess represents a structured success response for APIs.
type HTTPSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HandleError(ctx *fiber.Ctx, err error, statusCode int, message string) error {
	if err != nil {
		log.Println(err)
	}
	return WriteJSON(ctx, statusCode, nil, message)
}
