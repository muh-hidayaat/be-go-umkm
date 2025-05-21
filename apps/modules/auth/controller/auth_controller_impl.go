package controller

import (
	"be-go-umkm/apps/domain"
	"be-go-umkm/apps/helpers"
	"be-go-umkm/apps/modules/auth/request"
	"be-go-umkm/apps/modules/auth/service"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type UserControllerImpl struct {
	service     service.UserService
	redisClient *redis.Client
}

func NewUserController(service service.UserService, redisClient *redis.Client) UserController {
	return &UserControllerImpl{
		service:     service,
		redisClient: redisClient,
	}
}

func (c *UserControllerImpl) FindByID(ctx *fiber.Ctx) error {

	userUUID, err := helpers.ExtractUserID(ctx)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid user ID format")
	}
	fmt.Println("User ID from context:", userUUID)
	user, err := c.service.FindByID(ctx.Context(), userUUID)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusNotFound, "User not found")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, user, "User fetched successfully")
}

func (c *UserControllerImpl) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid user ID")
	}
	var user domain.User
	if err := ctx.BodyParser(&user); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}
	user.ID = id
	updatedUser, err := c.service.Update(ctx.Context(), user)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to update user")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, updatedUser, "User updated successfully")
}

func (c *UserControllerImpl) Register(ctx *fiber.Ctx) error {

	var req request.RegisterUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid input")
	}

	if err := req.Validate(); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, err.Error())
	}

	// Check if email already registered
	existingUser, err := c.service.FindByEmail(ctx.Context(), req.Email)
	if err == nil && existingUser.ID != uuid.Nil {
		return helpers.HandleError(ctx, err, fiber.StatusConflict, "Email already registered")
	}

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	registeredUser, err := c.service.Create(ctx.Context(), user)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to create user")
	}

	return helpers.WriteJSON(ctx, fiber.StatusCreated, registeredUser, "User registered successfully")
}

func (c *UserControllerImpl) Login(ctx *fiber.Ctx) error {
	request := new(request.UserLoginRequest)

	if err := ctx.BodyParser(request); err != nil {
		fmt.Println(err)
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid input")
	}

	if err := request.Validate(); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, err.Error())
	}
	user, err := c.service.FindByEmail(ctx.Context(), request.Email)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusUnauthorized, "Invalid credentials")
	}

	// Compare password
	err = c.service.ComparePassword(user.Password, request.Password)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusUnauthorized, "Invalid credentials")
	}

	// Generate auth token (JWT)
	token, err := c.service.GenerateAuthToken(user)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to generate auth token")
	}

	err = c.redisClient.Set(context.Background(), user.ID.String(), token, time.Hour*24*7).Err()
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to store refresh token")
	}

	return helpers.WriteJSON(ctx, fiber.StatusOK, map[string]interface{}{"token": token, "user": user}, "Login successful")
}

// func (c *UserControllerImpl) ForgotPassword(ctx *fiber.Ctx) error {
// 	var request struct {
// 		Email string `json:"email"`
// 	}

// 	if err := ctx.BodyParser(&request); err != nil {
// 		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
// 	}

// 	user, err := c.service.FindByEmail(ctx.Context(), request.Email)
// 	if err != nil {
// 		return helpers.HandleError(ctx, err, fiber.StatusNotFound, "User not found")
// 	}

// 	// Send password reset email
// 	err = c.service.SendPasswordResetEmail(user)
// 	if err != nil {
// 		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to send reset password email")
// 	}

// 	return helpers.WriteJSON(ctx, fiber.StatusOK, nil, "Password reset email sent")
// }

// func (c *UserControllerImpl) VerifyEmail(ctx *fiber.Ctx) error {
// 	token := ctx.Params("otp")
// 	email := ctx.Query("email", "")

// 	// Verify email by token
// 	err := c.service.VerifyEmail(ctx.Context(), token, email)
// 	if err != nil {
// 		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, err.Error())
// 	}

// 	return helpers.WriteJSON(ctx, fiber.StatusOK, nil, "Email verified successfully")
// }
