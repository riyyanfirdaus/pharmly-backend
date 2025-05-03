package server

import (
	"pharmly-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, handlers *RoutesOpts) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/register", middleware.ValidateRegisterRequest, handlers.AuthHandler.Register)
	auth.Post("/login", middleware.ValidateLoginRequest, handlers.AuthHandler.Login)
}
