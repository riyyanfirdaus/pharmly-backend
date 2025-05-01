package config

import "github.com/gofiber/fiber/v2"

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		AppName: "pharmly API",
	}
}

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"status":  "error",
		"message": err.Error(),
	})
}
