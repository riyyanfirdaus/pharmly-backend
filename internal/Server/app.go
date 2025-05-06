package server

import (
	"context"
	"log"
	"os"
	"pharmly-backend/config"
	"pharmly-backend/internal/database"
	"pharmly-backend/internal/handler"
	"pharmly-backend/internal/middleware"
	"pharmly-backend/internal/repository"
	"pharmly-backend/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type App struct {
	FiberApp *fiber.App
	DB       *database.PostgresDB
}

type RoutesOpts struct {
	AuthHandler     *handler.AuthHandler
	UserHandler     *handler.UserHandler
	CategoryHandler *handler.CategoryHandler
	ProductHandler  *handler.ProductHandler
	SupplierHandler *handler.SupplierHandler
}

func NewApp() (*App, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.NewPostgresDB()
	if err != nil {
		return nil, err
	}

	fiberConfig := config.NewFiberConfig()
	fiberConfig.ErrorHandler = middleware.ErrorHandler()
	fiberApp := fiber.New(fiberConfig)

	return &App{
		FiberApp: fiberApp,
		DB:       db,
	}, nil
}

func (a *App) Initialize() error {
	userRepo := repository.NewUserRepository(a.DB.Conn)
	categoryRepo := repository.NewCategoryRepository(a.DB.Conn)
	productRepo := repository.NewProductRepository(a.DB.Conn)
	supplierRepo := repository.NewSupplierRepository(a.DB.Conn)

	authUsecase := usecase.NewAuthUsecase(userRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	productUsecase := usecase.NewProductusecase(productRepo)
	supplierUsecase := usecase.NewSupplierUsecase(supplierRepo)

	authHandler := handler.NewAuthHandler(authUsecase)
	userHandler := handler.NewUserHandler(userUsecase)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	productHandler := handler.NewProductHandler(productUsecase)
	supplierHandler := handler.NewSupplierHandler(supplierUsecase)

	SetupRouter(a.FiberApp, &RoutesOpts{
		AuthHandler:     authHandler,
		UserHandler:     userHandler,
		CategoryHandler: categoryHandler,
		ProductHandler:  productHandler,
		SupplierHandler: supplierHandler,
	})

	return nil
}

func (a *App) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return a.FiberApp.Listen(":" + port)
}

func (a *App) Shutdown() error {
	return a.DB.Close(context.Background())
}
