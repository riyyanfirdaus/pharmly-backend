package handler

import (
	"pharmly-backend/internal/dto"
	"pharmly-backend/internal/logger"
	"pharmly-backend/internal/middleware"
	"pharmly-backend/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(usecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.UserRequest

	if err := c.BodyParser(&req); err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Interface("body", c.Body()).
			Msg("Failed to parse request body")
		return err
	}

	if err := middleware.Validate.Struct(&req); err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Interface("errors", middleware.GetValidationErrors(err)).
			Msg("Validation failed")
		return err
	}

	response, err := h.usecase.Register(c.Context(), &req)
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Msg("Failed to register user")
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User registered successfully",
		"data":    response,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Interface("body", c.Body()).
			Msg("Failed to parse request body")
		return err
	}

	if err := middleware.Validate.Struct(&req); err != nil {
		return err
	}

	response, err := h.usecase.Login(c.Context(), &req)
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Msg("Failed to login")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login successful",
		"data":    response,
	})
}
