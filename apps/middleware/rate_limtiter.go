package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        1000,            // Maksimum 10 permintaan
		Expiration: 1 * time.Minute, // dalam jangka waktu 1 menit
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Pembatasan berdasarkan IP address
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests. Please try again later.",
			})
		},
	})
}
