package controller

import (
	"be-go-umkm/apps/domain"
	"be-go-umkm/apps/helpers"
	"be-go-umkm/apps/modules/account/request"
	"be-go-umkm/apps/modules/account/service"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AccountControllerImpl struct {
	service    service.AccountService
	s3Client   *s3.S3
	bucketName string
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
	return helpers.WriteJSON(ctx, fiber.StatusOK, subcategories, "Subcategories fetched successfully")
}

func (c *AccountControllerImpl) FindByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid account ID")
	}

	account, err := c.service.FindByID(ctx.Context(), id)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusNotFound, "Subcategory not found")
	}

	return helpers.WriteJSON(ctx, fiber.StatusOK, account, "Subcategory fetched successfully")
}

func (c *AccountControllerImpl) Create(ctx *fiber.Ctx) error {
	var req request.AccountCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}

	if err := req.Validate(); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, err.Error())
	}

	account := domain.Account{
		Name:    req.Name,
		Type:    req.Type,
		Balance: req.Balance,
	}

	createdUser, err := c.service.Create(ctx.Context(), account)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to create user")
	}
	return helpers.WriteJSON(ctx, fiber.StatusCreated, createdUser, "User created successfully")
}
func (c *AccountControllerImpl) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid user ID")
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
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to update user")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, updatedAccount, "User updated successfully")
}

func (c *AccountControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid account ID")
	}

	if err := c.service.Delete(ctx.Context(), id); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to delete account")
	}

	return helpers.WriteJSON(ctx, fiber.StatusOK, nil, "Subcategory deleted successfully")
}
