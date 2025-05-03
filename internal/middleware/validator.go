package middleware

import (
	"pharmly-backend/internal/dto"
	"pharmly-backend/internal/logger"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

func ValidateRegisterRequest(c *fiber.Ctx) error {
	var req dto.UserRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Error().Err(err).Msg("Failed to parse registration request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	if req.Username == "" {
		logger.Error().Msg("Username is required")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Username is required",
		})
	}

	if len(req.Username) < 3 || len(req.Username) > 20 {
		logger.Error().Str("username", req.Username).Msg("Username length must be between 3 and 20 characters")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Username length must be between 3 and 20 characters",
		})
	}

	if req.Email == "" {
		logger.Error().Msg("Email is required")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Email is required",
		})
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(req.Email) {
		logger.Error().Str("email", req.Email).Msg("Invalid email format")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid email format",
		})
	}

	if req.Password == "" {
		logger.Error().Msg("Password is required")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Password is required",
		})
	}

	if len(req.Password) < 8 {
		logger.Error().Msg("Password must be at least 8 characters long")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Password must be at least 8 characters long",
		})
	}

	if req.Role == "" {
		req.Role = "cashier"
	} else if req.Role != "cashier" && req.Role != "pharmacist" && req.Role != "admin" {
		logger.Error().Str("role", req.Role).Msg("Invalid role")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Role must be either 'cashier' or 'pharmacist' or 'admin'",
		})
	}

	c.Locals("validatedRequest", &req)
	return c.Next()
}

func ValidateLoginRequest(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Error().Err(err).Msg("Failed to parse login request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	if req.Email == "" {
		logger.Error().Msg("Email is required")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Email is required",
		})
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(req.Email) {
		logger.Error().Str("email", req.Email).Msg("Invalid email format")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid email format",
		})
	}

	if req.Password == "" {
		logger.Error().Msg("Password is required")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Password is required",
		})
	}

	c.Locals("validatedRequest", &req)
	return c.Next()
}
