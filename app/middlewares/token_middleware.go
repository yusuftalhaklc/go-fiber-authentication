package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/utils"
)

func TokenMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Cookies("access_token")
		_, err := utils.VerifyToken(tokenString)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": err.Error()})
		}

		return c.Next()
	}
}
