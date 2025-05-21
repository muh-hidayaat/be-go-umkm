package controller

import (
	"be-go-umkm/apps/domain"
	"be-go-umkm/apps/helpers"
	"be-go-umkm/apps/modules/users/request"
	"be-go-umkm/apps/modules/users/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserControllerImpl struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		service: service,
	}
}

func (c *UserControllerImpl) FindAll(ctx *fiber.Ctx) error {
	user, err := c.service.FindAll(ctx.Context())
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to fetch user")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, user, "user fetched successfully")
}

func (c *UserControllerImpl) FindByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid user ID")
	}

	user, err := c.service.FindByID(ctx.Context(), id)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusNotFound, "user not found")
	}

	return helpers.WriteJSON(ctx, fiber.StatusOK, user, "User fetched successfully")
}

func (c *UserControllerImpl) Create(ctx *fiber.Ctx) error {
	var req request.CreateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}

	if err := req.Validate(); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, err.Error())
	}

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		// RoleID:   req.RoleID, // RoleID as uuid.UUID
	}

	createdUser, err := c.service.Create(ctx.Context(), user)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to create user")
	}
	return helpers.WriteJSON(ctx, fiber.StatusCreated, createdUser, "User created successfully")
}

func (c *UserControllerImpl) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid user ID")
	}

	var req request.UpdateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}

	if err := req.Validate(); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, err.Error())
	}

	user := domain.User{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	}

	updatedUser, err := c.service.Update(ctx.Context(), user)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to update user")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, updatedUser, "User updated successfully")
}

func (c *UserControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid user ID")
	}
	if err := c.service.Delete(ctx.Context(), id); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to delete user")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, nil, "User deleted successfully")
}

func (c *UserControllerImpl) ChangePassword(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid user ID")
	}

	var req request.UpdatePasswordRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}

	if err := req.Validate(); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, err.Error())
	}

	if err := c.service.ChangePassword(ctx.Context(), id, req.PasswordOld, req.Password, req.ConfirmPassword); err != nil {
		if err.Error() == "incorrect old password" {
			return helpers.HandleError(ctx, err, fiber.StatusUnauthorized, "Incorrect old password")
		}
		if err.Error() == "new password and confirm password do not match" {
			return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "New password and confirm password do not match")
		}
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to update password")
	}

	return helpers.WriteJSON(ctx, fiber.StatusOK, nil, "Password updated successfully")
}
