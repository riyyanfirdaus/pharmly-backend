package server

import (
	"context"
	"log"
	"os"
	"pharmly-backend/config"
	"pharmly-backend/internal/database"
	"pharmly-backend/internal/handler"
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
	AuthHandler *handler.AuthHandler
}

func NewApp() (*App, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.NewPostgresDB()
	if err != nil {
		return nil, err
	}

	fiberApp := fiber.New(config.NewFiberConfig())

	return &App{
		FiberApp: fiberApp,
		DB:       db,
	}, nil
}

func (a *App) Initialize() error {
	userRepo := repository.NewUserRepository(a.DB.Conn)

	authUsecase := usecase.NewAuthUsecase(userRepo)

	authHandler := handler.NewAuthHandler(authUsecase)

	SetupRouter(a.FiberApp, &RoutesOpts{
		AuthHandler: authHandler,
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
