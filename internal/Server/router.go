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

	v1.Use(middleware.AuthMiddleware())
	users := v1.Group("/users")
	users.Get("/", handlers.UserHandler.GetUsers)

	categories := v1.Group("/categories")
	categories.Get("/", handlers.CategoryHandler.GetCategories)

	products := v1.Group("/products")
	products.Post("/", handlers.ProductHandler.AddProduct)
	products.Get("/:id", handlers.ProductHandler.GetProductByID)
	products.Get("/", handlers.ProductHandler.GetProducts)
	products.Put("/", handlers.ProductHandler.UpdateProduct)
	products.Delete("/", handlers.ProductHandler.DeleteProduct)

	suppliers := v1.Group("/suppliers")
	suppliers.Get("/", handlers.SupplierHandler.GetSuppliers)
}
