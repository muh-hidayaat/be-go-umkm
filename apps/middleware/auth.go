package middleware

import (
	"be-go-umkm/apps/helpers"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

// AuthMiddleware untuk autentikasi berbasis userID dan role (untuk login)
func AuthMiddleware(rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil token dari header
		tokenString, err := getTokenFromHeader(c)
		if err != nil {
			return unauthorizedResponse(c, "Invalid or missing token")
		}

		// Verifikasi dan parsing token JWT
		claims, err := parseAndValidateToken(tokenString)
		if err != nil {
			log.Println("Token validation error:", err)
			return unauthorizedResponse(c, "Invalid or expired token")
		}

		// Ekstrak userID dari token
		userID, err := extractUserClaims(claims)
		if err != nil {
			log.Println("Claim extraction error:", err)
			return unauthorizedResponse(c, "Invalid token payload")
		}

		// Validasi token di Redis
		if err := validateTokenInRedis(rdb, userID, tokenString); err != nil {
			log.Println("Token Redis validation error:", err)
			return unauthorizedResponse(c, "Unauthorized access")
		}

		// Periksa apakah role pengguna diizinkan
		// if !isAuthorizedRole(allowedRoles, role) {
		// 	// fmt.Println(allowedRoles)
		// 	// fmt.Println(role)
		// 	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access Denied"})
		// }

		// Simpan userID di konteks request
		c.Locals("userID", userID)
		return c.Next()
	}
}

// ForgotPasswordMiddleware untuk validasi email (tanpa role)
func ForgotPasswordMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil token dari header
		tokenString, err := getTokenFromHeader(c)
		if err != nil {
			return unauthorizedResponse(c, "Invalid or missing token")
		}

		// Verifikasi token JWT
		claims, err := parseAndValidateToken(tokenString)
		if err != nil {
			log.Println("Token validation error:", err)
			return unauthorizedResponse(c, "Invalid or expired token")
		}

		// Ekstrak email dari token
		email, err := extractEmailClaim(claims)
		if err != nil {
			log.Println("Claim extraction error:", err)
			return unauthorizedResponse(c, "Invalid token payload")
		}

		// Simpan email di konteks request
		c.Locals("email", email)
		return c.Next()
	}
}

// getTokenFromHeader mengambil JWT dari header Authorization
func getTokenFromHeader(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("missing or invalid token format")
	}
	return strings.TrimPrefix(authHeader, "Bearer "), nil
}

// parseAndValidateToken memverifikasi dan mem-parsing token JWT
func parseAndValidateToken(tokenString string) (jwt.MapClaims, error) {
	return helpers.VerifyJWT(tokenString)
}

// extractUserClaims mengambil userID dan role dari JWT claims
func extractUserClaims(claims jwt.MapClaims) (string, error) {
	userID, ok := claims["userID"].(string)
	if !ok {
		return "", fmt.Errorf("invalid userID claim in token")
	}

	// customerID, _ := claims["customerID"].(string)

	// role, ok := claims["role"].(string)
	// if !ok {
	// 	return "", "", "", fmt.Errorf("invalid role claim in token")
	// }

	return userID, nil
}

// extractEmailClaim mengambil email dari JWT claims
func extractEmailClaim(claims jwt.MapClaims) (string, error) {
	email, ok := claims["email"].(string)
	if !ok {
		return "", fmt.Errorf("invalid email claim in token")
	}
	return email, nil
}

// validateTokenInRedis memeriksa apakah token tersimpan di Redis (opsional)
func validateTokenInRedis(rdb *redis.Client, userID, tokenString string) error {

	storedToken, err := rdb.Get(context.Background(), userID).Result()

	if err == redis.Nil || storedToken != tokenString {
		return fmt.Errorf("unauthorized")
	} else if err != nil {
		return fmt.Errorf("internal server error")
	}
	return nil
}

// isAuthorizedRole memeriksa apakah role pengguna ada dalam daftar role yang diperbolehkan
// func isAuthorizedRole(allowedRoles []string, userRole string) bool {
// 	for _, role := range allowedRoles {
// 		if role == userRole {
// 			return true
// 		}
// 	}
// 	return false
// }

// unauthorizedResponse mengembalikan respons 401 Unauthorized dengan pesan yang lebih jelas
func unauthorizedResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": message})
}
