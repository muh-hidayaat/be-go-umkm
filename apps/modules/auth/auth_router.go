package auth

import (
	"be-go-umkm/apps/middleware"
	"be-go-umkm/apps/modules/auth/controller"
	"be-go-umkm/apps/modules/auth/repository"
	"be-go-umkm/apps/modules/auth/service"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// @Summary Auth Routes
// @Description API endpoints for authentication and authorization
// @Tags Auth
// @Accept json
// @Produce json
func Router(app fiber.Router, db *gorm.DB, rdb *redis.Client) {
	repo := repository.NewUserRepository()
	// repoCustomer := repoCustomer.NewCustomerRepository()
	svc := service.NewUserService(repo, db)
	ctrl := controller.NewUserController(svc, rdb)

	// @Summary Register a new user
	// @Description Register a new user with email and password
	// @Tags Auth
	// @Accept json
	// @Produce json
	// @Router /register [post]
	app.Post("/register", ctrl.Register) // Register a new user

	// @Summary Login a user
	// @Description Login a user and return a JWT token
	// @Tags Auth
	// @Accept json
	// @Produce json
	// @Router /login [post]
	app.Post("/login", ctrl.Login) // Login a user

	r := app.Group("/auth")

	// @Summary Fetch authenticated user details
	// @Description Get details of the currently authenticated user
	// @Tags Auth
	// @Accept json
	// @Produce json
	// @Router /auth/user [get]
	r.Get("/user", middleware.AuthMiddleware(rdb), ctrl.FindByID) // Fetch authenticated user details

	// @Summary Verify user email
	// @Description Verify user email using OTP
	// @Tags Auth
	// @Accept json
	// @Produce json
	// @Router /auth/email/verify/{otp} [post]
	// r.Post("/email/verify/:otp", ctrl.VerifyEmail) // Verify user email

}
