package middleware

import (
	"pharmly-backend/internal/logger"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	Validate.RegisterValidation("password", validatePassword)
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 8 {
		return false
	}

	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false
	}

	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return false
	}

	if !strings.ContainsAny(password, "0123456789") {
		return false
	}

	if !strings.ContainsAny(password, "!@#$%^&*()_+-=[]{}|;:,.<>?") {
		return false
	}

	return true
}

func GetValidationErrors(err error) map[string]string {
	errs := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := strings.ToLower(e.Field())
			switch e.Tag() {
			case "required":
				errs[field] = "This field is required"
			case "email":
				errs[field] = "Invalid email format"
			case "password":
				errs[field] = "Password must be at least 8 characters long and contain uppercase, lowercase, number, and special character"
			case "len":
				errs[field] = "Invalid length"
			case "gte":
				errs[field] = "Value must be greater than or equal to " + e.Param()
			case "lte":
				errs[field] = "Value must be less than or equal to " + e.Param()
			default:
				errs[field] = "Invalid value"
			}
		}
	}

	return errs
}

func ErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			logger.Error().
				Err(err).
				Interface("errors", GetValidationErrors(validationErrors)).
				Str("path", c.Path()).
				Str("method", c.Method()).
				Msg("Validation failed")

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Validation failed",
				"errors":  GetValidationErrors(validationErrors),
			})
		}

		if e, ok := err.(*fiber.Error); ok && e.Code == fiber.StatusBadRequest {
			logger.Error().
				Err(err).
				Str("path", c.Path()).
				Str("method", c.Method()).
				Msg("Invalid request body")

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid request body",
				"errors": map[string]string{
					"body": "Request body is invalid or malformed",
				},
			})
		}

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			logger.Error().
				Err(err).
				Int("status_code", code).
				Str("path", c.Path()).
				Str("method", c.Method()).
				Msg("Request failed")

			return c.Status(code).JSON(fiber.Map{
				"status":  "error",
				"message": e.Message,
			})
		}

		logger.Error().
			Err(err).
			Int("status_code", code).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Msg("Request failed")

		return c.Status(code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
}

func ValidateRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() == "GET" || c.Method() == "DELETE" {
			return c.Next()
		}

		var body map[string]interface{}
		if err := c.BodyParser(&body); err != nil {
			logger.Error().
				Err(err).
				Str("path", c.Path()).
				Str("method", c.Method()).
				Msg("Failed to parse request body")
			return err
		}

		if len(body) == 0 {
			logger.Error().
				Str("path", c.Path()).
				Str("method", c.Method()).
				Msg("Request body cannot be empty")
			return fiber.NewError(fiber.StatusBadRequest, "Request body cannot be empty")
		}

		return c.Next()
	}
}
