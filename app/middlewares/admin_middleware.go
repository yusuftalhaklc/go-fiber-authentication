package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/models"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/repositories"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/utils"
)

// AdminMiddleware is a middleware function that is used to authenticate and authorize an admin user.
// It retrieves the user data from the repository based on the access token stored in the cookie.
// If the user is authenticated and has an admin role, the middleware allows the request to proceed.
// Otherwise, it returns an error response indicating unauthorized access.
func AdminMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the user data from the repository based on the access token
		var foundUser *models.GetResponse
		userRepository := repositories.NewUserRepository()
		tokenString := c.Cookies("access_token")

		// Verify the token and retrieve the claims
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
		}

		// Retrieve the user from the repository based on the user's email
		userEmail := claims["email"].(string)
		foundUser, err = userRepository.GetUser(userEmail)
		if err != nil {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
		}

		// Check if the user has an admin role
		if !(models.IsAdminValid(foundUser.UserRole)) {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Unauthorized"})
		}

		// Allow the request to proceed
		return c.Next()
	}
}
