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
		// Retrieve the access token from the cookie
		tokenString := c.Cookies("access_token")

		// Verify the token
		_, err := utils.VerifyToken(tokenString)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": err.Error()})
		}

		// Allow the request to proceed
		return c.Next()
	}
}
