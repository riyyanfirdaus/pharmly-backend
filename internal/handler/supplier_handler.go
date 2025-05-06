package handler

import (
	"pharmly-backend/internal/logger"
	"pharmly-backend/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type SupplierHandler struct {
	usecase usecase.SupplierUsecase
}

func NewSupplierHandler(usecase usecase.SupplierUsecase) *SupplierHandler {
	return &SupplierHandler{usecase: usecase}
}

func (h *SupplierHandler) GetSuppliers(c *fiber.Ctx) error {
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

	suppliers, pagination, err := h.usecase.GetAllSuppliers(c.Context(), page, pageSize)
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Msg("Failed to get suppliers")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"message":    "Suppliers retrieved successfully",
		"data":       suppliers,
		"pagination": pagination,
	})
}
