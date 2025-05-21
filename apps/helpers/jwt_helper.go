package helpers

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func GenerateJWTByEmail(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	})

	return token.SignedString([]byte(viper.GetString("JWT_SECRET_KEY")))
}

func ExtractEmailFromToken(ctx *fiber.Ctx) (string, error) {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}

	tokenStr := authHeader[len("Bearer "):]
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("invalid token format")
	}

	return email, nil
}

// GenerateJWT generates a JWT for the given user ID.
func GenerateJWT(userID uuid.UUID, customerID *uuid.UUID, role string) (string, error) {
	SecretKey := viper.GetString("JWT_SECRET_KEY")
	if SecretKey == "" {
		return "", fmt.Errorf("JWT_SECRET_KEY is not set in the configuration")
	}

	claims := jwt.MapClaims{
		"userID": userID.String(),
		"role":   role,
		"exp":    time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	if customerID != nil {
		claims["customerID"] = customerID.String()
	} else {
		claims["customerID"] = nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SecretKey))
}

// VerifyJWT verifies the given JWT and returns the claims if valid.
func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	SecretKey := viper.GetString("JWT_SECRET_KEY")
	if SecretKey == "" {
		return nil, fmt.Errorf("JWT_SECRET_KEY is not set in the configuration")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	exp, ok := claims["exp"].(float64)
	if !ok || time.Now().Unix() > int64(exp) {
		return nil, fmt.Errorf("token is expired")
	}

	return claims, nil
}
