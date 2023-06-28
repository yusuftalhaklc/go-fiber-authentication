package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/models"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/repositories"
)

func Signup(c *fiber.Ctx) error {
	var user models.User

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	userRepository := repositories.NewUserRepository()
	err = userRepository.Create(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})
}

func Login(c *fiber.Ctx) error {
	var user models.User

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	userRepository := repositories.NewUserRepository()
	err = userRepository.Login(&user)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	loginResponse := models.LoginResponse{Email: user.Email, Token: user.Token}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Successfully login", "data": loginResponse})
}

func Logout(c *fiber.Ctx) error {
	authToken := c.Get("Authorization")
	userRepository := repositories.NewUserRepository()

	err := userRepository.Logout(&authToken)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Successfully logout"})
}

func Delete(c *fiber.Ctx) error {
	authToken := c.Get("Authorization")
	userRepository := repositories.NewUserRepository()

	err := userRepository.Delete(&authToken)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "success", "message": "Successfully deleted"})
}
