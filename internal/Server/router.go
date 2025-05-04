package server

import (
	"pharmly-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, handlers *RoutesOpts) {
	app.Use(middleware.ValidateRequest())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/register", handlers.AuthHandler.Register)
	auth.Post("/login", handlers.AuthHandler.Login)

	users := v1.Group("/users")
	users.Use(middleware.AuthMiddleware())
	users.Get("/", handlers.UserHandler.GetUsers)
}
