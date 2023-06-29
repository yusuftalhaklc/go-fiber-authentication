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

func Update(c *fiber.Ctx) error {
	authToken := c.Get("Authorization")
	var user models.User

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	userRepository := repositories.NewUserRepository()
	err = userRepository.Update(&user, &authToken)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Successfully updated"})
}
