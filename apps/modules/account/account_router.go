package account

import (
	"be-go-umkm/apps/modules/account/controller"
	"be-go-umkm/apps/modules/account/repository"
	"be-go-umkm/apps/modules/account/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app fiber.Router, db *gorm.DB) {
	repo := repository.NewAccountRepository()
	svc := service.NewAccountService(repo, db)
	ctrl := controller.NewAccountController(svc)

	r := app.Group("/subcategories")
	r.Get("/", ctrl.FindAll)
	r.Get("/:id", ctrl.FindByID)
	r.Post("/", ctrl.Create)
	r.Put("/:id", ctrl.Update)
	r.Delete("/:id", ctrl.Delete)
}
