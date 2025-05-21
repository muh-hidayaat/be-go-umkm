package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// ExtractUserID extracts and validates the user ID from the request context.
func ExtractUserID(ctx *fiber.Ctx) (uuid.UUID, error) {
	userID := ctx.Locals("userID")
	if userID == nil {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "User ID not found in token")
	}

	userStr, ok := userID.(string)
	if !ok {
		return uuid.Nil, fiber.NewError(fiber.StatusBadRequest, "Invalid user ID format")
	}

	parsedID, err := uuid.Parse(userStr)
	if err != nil {
		return uuid.Nil, fiber.NewError(fiber.StatusBadRequest, "Invalid user ID UUID")
	}

	return parsedID, nil
}

// ExtractCustomerID extracts and validates the customer ID from the request context.
func ExtractCustomerID(ctx *fiber.Ctx) (uuid.UUID, error) {
	customerID := ctx.Locals("customerID")
	if customerID == nil {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "Customer ID not found in token")
	}

	customerStr, ok := customerID.(string)
	if !ok {
		return uuid.Nil, fiber.NewError(fiber.StatusBadRequest, "Invalid customer ID format")
	}

	parsedID, err := uuid.Parse(customerStr)
	if err != nil {
		return uuid.Nil, fiber.NewError(fiber.StatusBadRequest, "Invalid customer ID UUID")
	}

	return parsedID, nil
}
