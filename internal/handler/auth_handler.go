package handler

import (
	"pharmly-backend/internal/dto"
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
	req := c.Locals("validatedRequest").(*dto.UserRequest)

	response, err := h.usecase.Register(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   response,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := c.Locals("validatedRequest").(*dto.LoginRequest)

	response, err := h.usecase.Login(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "succcess",
		"data":   response,
	})
}
