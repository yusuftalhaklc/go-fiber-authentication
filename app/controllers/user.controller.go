package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/models"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/repositories"
)

func GetUser(c *fiber.Ctx) error {
	authToken := c.Get("Authorization")
	var foundUser *models.GetResponse
	userRepository := repositories.NewUserRepository()

	foundUser, err := userRepository.GetUser(&authToken)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Successfully found", "data": foundUser})
}
