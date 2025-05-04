package handler

import (
	"pharmly-backend/internal/logger"
	"pharmly-backend/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Str("page", c.Query("page")).
			Msg("Invalid page number")
		return err
	}

	pageSize, err := strconv.Atoi(c.Query("page_size", "10"))
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Str("page_size", c.Query("page_size")).
			Msg("Invalid page size")
		return err
	}

	users, pagination, err := h.usecase.GetAllUsers(c.Context(), page, pageSize)
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Msg("Failed to get users")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"message":    "Users retrieved successfully",
		"data":       users,
		"pagination": pagination,
	})
}
