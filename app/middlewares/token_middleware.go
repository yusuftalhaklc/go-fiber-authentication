package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/utils"
)

// TokenMiddleware is a middleware function that is used to verify the access token provided in the request.
// It retrieves the access token from the cookie and verifies its validity using the VerifyToken function.
// If the token is invalid or expired, it returns an error response indicating unauthorized access.
// Otherwise, it allows the request to proceed.
func TokenMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the JWT token from the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" || len(authHeader) < 7 {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "invalid token"})
		}
		tokenString := authHeader[7:] // Remove the "Bearer " prefix

		// Verify the token
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": err.Error()})
		}
		c.Locals("claims", claims)
		// Allow the request to proceed
		return c.Next()
	}
}
