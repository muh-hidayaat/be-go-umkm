package router

import (
	"be-go-umkm/apps/middleware"
	account "be-go-umkm/apps/modules/account"
	auth "be-go-umkm/apps/modules/auth"
	user "be-go-umkm/apps/modules/users"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, rdb *redis.Client, s3Client *s3.S3, bucketName string) {
	app.Use(middleware.RateLimiter())

	apiRoutes := app.Group("/api/v1")

	auth.Router(apiRoutes, db, rdb)
	user.Router(apiRoutes, db, rdb)
	account.Router(apiRoutes, db, rdb)
	// auouter(apiRoutes, db, rdb, s3Client, bucketName)
}
