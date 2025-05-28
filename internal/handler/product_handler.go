package handler

import (
	"pharmly-backend/internal/dto"
	"pharmly-backend/internal/logger"
	"pharmly-backend/internal/middleware"
	"pharmly-backend/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	usecase usecase.ProductUsecase
}

func NewProductHandler(usecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: usecase}
}

func (h *ProductHandler) AddProduct(c *fiber.Ctx) error {
	var req dto.ProductRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Interface("body", c.Body()).
			Msg("Failed to parse request body")
		return err
	}

	if err := middleware.Validate.Struct(&req); err != nil {
		return err
	}

	response, err := h.usecase.CreateProduct(c.Context(), &req)
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Msg("Failed to add product")
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "product added successfully",
		"data":    response,
	})
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Str("id", c.Params("id")).
			Msg("Invalid product ID")
		return nil
	}

	product, err := h.usecase.GetProductByID(c.Context(), id)
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Int64("id", id).
			Msg("Failed to get product")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Product retrieved successfully",
		"data":    product,
	})
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
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

	products, pagination, err := h.usecase.GetAllProducts(c.Context(), page, pageSize)
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Msg("Failed to get products")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"message":    "Products retrieved successfully",
		"data":       products,
		"pagination": pagination,
	})
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Str("id", c.Params("id")).
			Msg("Invalid product ID")
		return nil
	}

	var req dto.ProductRequest
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

	product, err := h.usecase.UpdateProduct(c.Context(), id, &req)
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Int64("id", id).
			Msg("Failed to update product")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Product updated successfully",
		"data":    product,
	})
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Str("id", c.Params("id")).
			Msg("Invalid product ID")
		return nil
	}

	if err := h.usecase.DeleteProduct(c.Context(), id); err != nil {
		logger.Error().
			Err(err).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Int64("id", id).
			Msg("Failed to delete product")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Product deleted successfully",
	})
}
