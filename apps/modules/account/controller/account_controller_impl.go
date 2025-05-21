package controller

import (
	"be-go-umkm/apps/domain"
	"be-go-umkm/apps/helpers"
	"be-go-umkm/apps/modules/account/request"
	"be-go-umkm/apps/modules/account/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AccountControllerImpl struct {
	service service.AccountService
}

func NewAccountController(service service.AccountService) AccountController {
	return &AccountControllerImpl{
		service: service,
	}
}

func (c *AccountControllerImpl) FindAll(ctx *fiber.Ctx) error {
	subcategories, err := c.service.FindAll(ctx.Context())
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to fetch subcategories")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, subcategories, "Account fetched successfully")
}

func (c *AccountControllerImpl) FindByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid account ID")
	}

	account, err := c.service.FindByID(ctx.Context(), id)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusNotFound, "Account not found")
	}

	return helpers.WriteJSON(ctx, fiber.StatusOK, account, "Account fetched successfully")
}

func (c *AccountControllerImpl) Create(ctx *fiber.Ctx) error {
	var req request.AccountCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}

	if err := req.Validate(); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, err.Error())
	}

	userID, err := helpers.ExtractUserID(ctx)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid user ID format")
	}

	account := domain.Account{
		UserID:  userID,
		Name:    req.Name,
		Type:    req.Type,
		Balance: req.Balance,
	}

	createdAccount, err := c.service.Create(ctx.Context(), account)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to create account")
	}
	return helpers.WriteJSON(ctx, fiber.StatusCreated, createdAccount, "User created successfully")
}
func (c *AccountControllerImpl) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid account ID")
	}

	var req request.AccountUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}

	if err := req.Validate(); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, err.Error())
	}

	account := domain.Account{
		ID:      id,
		Name:    req.Name,
		Type:    req.Type,
		Balance: req.Balance,
	}

	updatedAccount, err := c.service.Update(ctx.Context(), account)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to update account")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, updatedAccount, "Account updated successfully")
}

func (c *AccountControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid account ID")
	}

	if err := c.service.Delete(ctx.Context(), id); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to delete account")
	}

	return helpers.WriteJSON(ctx, fiber.StatusOK, nil, "Account deleted successfully")
}
