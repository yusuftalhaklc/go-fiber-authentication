package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/handlers"
	middlewares "github.com/yusuftalhaklc/go-fiber-authentication/app/middlewares"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/user")

	middleware := middlewares.AuthMiddleware()

	v1.Post("/signup", handlers.Signup)
	v1.Post("/login", handlers.Login)

	v1.Put("/update", middleware, handlers.Update)
	v1.Get("/", middleware, handlers.GetUser)
	v1.Post("/logout", middleware, handlers.Logout)
	v1.Post("/delete", middleware, handlers.Delete)
}
