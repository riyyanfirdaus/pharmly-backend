package handler

import (
	"pharmly-backend/internal/logger"
	"pharmly-backend/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	usecase usecase.CategoryUsecase
}

func NewCategoryHandler(usecase usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{usecase: usecase}
}

func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "2"))
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Str("page", c.Query("page")).
			Msg("Invalid page number")
		return err
	}

	pageSize, err := strconv.Atoi(c.Query("page_size", "3"))
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Str("page_size", c.Query("page_size")).
			Msg("Invalid page size")
		return err
	}

	categories, pagination, err := h.usecase.GetAllCategories(c.Context(), page, pageSize)
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Msg("Failed to get categories")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"message":    "Categories retrieved successfully",
		"data":       categories,
		"pagination": pagination,
	})
}
