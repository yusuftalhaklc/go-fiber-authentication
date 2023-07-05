package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/models"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/repositories"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/utils"
)

func AdminMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var foundUser *models.GetResponse
		userRepository := repositories.NewUserRepository()
		tokenString := c.Cookies("access_token")

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
		}
		userEmail := claims["email"].(string)
		foundUser, err = userRepository.GetUser(userEmail)
		if err != nil {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
		}

		if !(models.IsAdminValid(foundUser.UserRole)) {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "unauthorized"})
		}
		return c.Next()
	}
}
