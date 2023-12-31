package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/config"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/router"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: config.Config("COOKIE_ENC_KEY"),
	}))

	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "not found endpoint"})
	})

	app.Listen(":8080")
}
