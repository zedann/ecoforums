package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zedann/ecoforum/server/internal/user"
)

func AuthRequire() fiber.Handler {

	return func(c *fiber.Ctx) error {

		tokenStr := c.Cookies("jwt")

		if tokenStr == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authentication required , please log in.",
			})
		}
		// Retrieve the secret key from environment variable
		secretKey := os.Getenv("SECRET_KEY")
		if secretKey == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Secret key not set",
			})
		}
		token, err := jwt.ParseWithClaims(tokenStr, &user.MyJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token, please log in again.",
			})
		}

		if claims, ok := token.Claims.(*user.MyJWTClaims); ok && token.Valid {
			if claims.ExpiresAt != nil && time.Until(claims.ExpiresAt.Time) < 0 {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
					"error": "Token has expired, please log in again.",
				})
			}

			c.Locals("user_id", claims.ID)
			c.Locals("username", claims.Username)
			return c.Next()
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

}
