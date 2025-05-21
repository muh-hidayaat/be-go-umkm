package user

import (
	"be-go-umkm/apps/modules/users/controller"
	"be-go-umkm/apps/modules/users/repository"
	"be-go-umkm/apps/modules/users/service"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// @Summary User Routes
// @Description API endpoints for managing users
// @Tags Users
// @Accept json
// @Produce json
func Router(app fiber.Router, db *gorm.DB, rdb *redis.Client) {
	repo := repository.NewUserRepository()
	svc := service.NewUserService(repo, db)
	ctrl := controller.NewUserController(svc)

	r := app.Group("/users") // Sesuaikan role sesuai kebutuhan

	r.Get("/", ctrl.FindAll)
	r.Get("/:id", ctrl.FindByID)
	r.Post("/", ctrl.Create)
	r.Put("/:id", ctrl.Update)
	r.Delete("/:id", ctrl.Delete)
	r.Put("/:id/change-password", ctrl.ChangePassword)
}
