package main

import (
	"be-go-umkm/apps/config"
	"be-go-umkm/apps/router"
	"fmt"
	"log"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	// Find and read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}

func setup() *fiber.App {
	db := config.DBConnect()
	if db == nil {
		log.Fatal("Failed to connect to database")
	}

	rdb := config.CreateRedisClient()
	if rdb == nil {
		log.Fatal("Failed to connect to Redis")
	}

	// _, err := config.ConnectToMongoDB()
	// if err != nil {
	// 	log.Fatalf("Error saat menghubungkan ke MongoDB: %v", err)
	// }

	app := fiber.New()

	// Initialize S3 client
	s3Config, err := config.InitS3()
	if err != nil {
		log.Printf("Error initializing S3 client: %s", err) // Error yang tidak menghentikan aplikasi
	}

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("Origin: %s\n", c.Get("Origin"))
		return c.Next()
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080,http://127.0.0.1:3005",
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Authorization",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		ExposeHeaders:    "Authorization", // Allow token barrier
	}))

	// Setup routes
	router.SetupRoutes(app, db, rdb, s3Config.Service, s3Config.BucketName)

	return app
}

func main() {
	fmt.Println("Application is starting...")

	app := setup()

	// Get APP_PORT from config and handle default case
	appPortStr := viper.GetString("APP_PORT")
	if appPortStr == "" {
		appPortStr = "8080" // Default port if not found in config
	}
	app_port, err := strconv.Atoi(appPortStr)
	if err != nil {
		log.Fatalf("Invalid APP_PORT value in config: %v", err)
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%d", app_port)))
}

// cek
