package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/models"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/repositories"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/utils"
)

func Signup(c *fiber.Ctx) error {
	user := new(models.User)

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	userRepository := repositories.NewUserRepository()
	user, err = userRepository.Create(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})
}

func Login(c *fiber.Ctx) error {
	user := new(models.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	userRepository := repositories.NewUserRepository()
	user, err = userRepository.Login(user)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	token, err := utils.CreateToken(user)
	if err != nil {
		return errors.New("token was not created")
	}

	cookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		HTTPOnly: true,
	}
	c.Cookie(cookie)

	loginResponse := models.LoginResponse{Email: user.Email, Token: token}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Successfully login", "data": loginResponse})
}

func Logout(c *fiber.Ctx) error {
	userRepository := repositories.NewUserRepository()
	tokenString := c.Cookies("access_token")

	claims, err := utils.VerifyToken(tokenString)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	invalidToken, err := utils.InvalidateToken(tokenString)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	cookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    invalidToken,
		HTTPOnly: true,
	}
	c.Cookie(cookie)
	userMail := claims["email"].(string)
	err = userRepository.Logout(userMail)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Successfully logout"})
}

func Delete(c *fiber.Ctx) error {
	userRepository := repositories.NewUserRepository()

	err := userRepository.Delete("")
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "success", "message": "Successfully deleted"})
}

func GetUser(c *fiber.Ctx) error {
	var foundUser *models.GetResponse
	userRepository := repositories.NewUserRepository()

	foundUser, err := userRepository.GetUser("")
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Successfully found", "data": foundUser})
}

func Update(c *fiber.Ctx) error {
	var user models.User

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	userRepository := repositories.NewUserRepository()
	err = userRepository.Update(&user)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Successfully updated"})
}
