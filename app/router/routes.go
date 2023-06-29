package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/user")

	v1.Post("/signup", controllers.Signup)
	v1.Post("/login", controllers.Login)
	v1.Put("/update", controllers.Update)
	v1.Get("/", controllers.GetUser)
	v1.Post("/logout", controllers.Logout)
	v1.Post("/delete", controllers.Delete)
}
