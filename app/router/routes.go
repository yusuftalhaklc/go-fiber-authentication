package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/handlers"
	middlewares "github.com/yusuftalhaklc/go-fiber-authentication/app/middlewares"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/user")

	tokenMiddleware := middlewares.TokenMiddleware()
	adminMiddleware := middlewares.AdminMiddleware()

	v1.Post("/signup", handlers.Signup)
	v1.Post("/login", handlers.Login)

	v1.Put("/update", tokenMiddleware, handlers.Update)
	v1.Get("/", tokenMiddleware, handlers.GetUser)
	v1.Post("/logout", tokenMiddleware, handlers.Logout)
	v1.Delete("/delete/:email", tokenMiddleware, adminMiddleware, handlers.Delete)
}
