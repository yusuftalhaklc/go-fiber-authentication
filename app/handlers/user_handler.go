package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/models"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/repositories"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/utils"
)

// Signup handles the signup process by parsing the user data from the request body,
// creating a new user in the repository, and returning the created user in the response.
// If there is an error during parsing or user creation, an error response is returned.
func Signup(c *fiber.Ctx) error {
	// Parse the user data from the request body
	user := new(models.User)
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	// Create a new user in the repository
	userRepository := repositories.NewUserRepository()
	user, err = userRepository.Create(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	// Return a success response with the created user data
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has been created", "data": user})
}

// Login handles the login process by parsing the user data from the request body,
// verifying the user's credentials in the repository, creating a token for the user,
// and setting the token as a cookie in the response. Finally, it returns a success response.
// If there is an error during parsing, user authentication, token creation, or setting the cookie, an error response is returned.
func Login(c *fiber.Ctx) error {
	// Parse the user data from the request body
	user := new(models.User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	// Verify the user's credentials in the repository
	userRepository := repositories.NewUserRepository()
	user, err = userRepository.Login(user)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	// Create a token for the user
	token, err := utils.CreateToken(user)
	if err != nil {
		return errors.New("token was not created")
	}

	// Set the token as a cookie in the response
	cookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		HTTPOnly: true,
	}
	c.Cookie(cookie)

	// Return a success response with the user's email and token
	loginResponse := models.LoginResponse{Email: user.Email, Token: token}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Successfully logged in", "data": loginResponse})
}

// Logout handles the logout process by invalidating the user's token,
// setting an invalidated token as a cookie in the response, and logging out the user in the repository.
// Finally, it returns a success response.
// If there is an error during token verification, token invalidation, cookie setting, or user logout, an error response is returned.
func Logout(c *fiber.Ctx) error {
	// Get the user's token from the cookie
	userRepository := repositories.NewUserRepository()
	tokenString := c.Cookies("access_token")

	// Verify the token and retrieve the claims
	claims, err := utils.VerifyToken(tokenString)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	// Invalidate the token
	invalidToken, err := utils.InvalidateToken(tokenString)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	// Set the invalidated token as a cookie in the response
	cookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    invalidToken,
		HTTPOnly: true,
	}
	c.Cookie(cookie)

	// Retrieve the user's email from the claims and log out the user in the repository
	userEmail := claims["email"].(string)
	err = userRepository.Logout(userEmail)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	// Return a success response
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Successfully logged out"})
}

// Delete handles the deletion of a user by email.
// It deletes the user from the repository and returns a success response.
// If there is an error during the deletion, an error response is returned.
func Delete(c *fiber.Ctx) error {
	// Get the email parameter from the request
	userRepository := repositories.NewUserRepository()
	emailParam := c.Params("email")

	// Delete the user from the repository
	err := userRepository.Delete(emailParam)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	// Return a success response
	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "success", "message": "Successfully deleted"})
}

// GetUser retrieves the user data for the authenticated user.
// It verifies the user's token, retrieves the user from the repository,
// and returns the user data in the response.
// If there is an error during token verification or retrieving the user, an error response is returned.
func GetUser(c *fiber.Ctx) error {
	// Create a variable to store the found user data
	var foundUser *models.GetResponse
	userRepository := repositories.NewUserRepository()

	// Get the user's token from the cookie
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

	// Return a success response with the found user data
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Successfully found", "data": foundUser})
}

// Update handles the update of a user's data.
// It parses the updated user data from the request body,
// updates the user in the repository, and returns a success response.
// If there is an error during parsing or updating the user data, an error response is returned.
func Update(c *fiber.Ctx) error {
	// Parse the updated user data from the request body
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	// Update the user in the repository
	userRepository := repositories.NewUserRepository()
	err = userRepository.Update(&user)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	// Return a success response
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Successfully updated"})
}
