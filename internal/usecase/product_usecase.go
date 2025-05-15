package usecase

import (
	"context"
	"pharmly-backend/internal/dto"
	"pharmly-backend/internal/entity"
	"pharmly-backend/internal/logger"
	"pharmly-backend/internal/repository"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req *dto.ProductRequest) (*dto.ProductResponse, error)
	GetAllProducts(ctx context.Context, page, pageSize int) ([]*entity.Product, *dto.PaginationResponse, error)
}

type productsUsecase struct {
	repo repository.ProductRepository
}

func NewProductusecase(repo repository.ProductRepository) ProductUsecase {
	return &productsUsecase{repo: repo}
}

func (u *productsUsecase) CreateProduct(ctx context.Context, req *dto.ProductRequest) (*dto.ProductResponse, error) {
	product := &entity.Product{
		Name:           req.Name,
		CategoryID:     req.CategoryID,
		GenericName:    req.GenericName,
		Description:    &req.Description,
		Price:          req.Price,
		Stock:          req.Stock,
		Unit:           req.Unit,
		ExpirationDate: req.ExpirationDate,
		Barcode:        req.Barcode,
		SupplierID:     req.SupplierID,
		MinStock:       req.MinStock,
		IsActive:       req.IsActive,
	}

	if err := u.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:             product.ID,
		Name:           product.Name,
		CategoryID:     product.CategoryID,
		GenericName:    product.GenericName,
		Description:    product.Description,
		Price:          product.Price,
		Stock:          product.Stock,
		Unit:           product.Unit,
		ExpirationDate: product.ExpirationDate,
		Barcode:        product.Barcode,
		SupplierID:     product.SupplierID,
		MinStock:       product.MinStock,
		IsActive:       product.IsActive,
	}, nil
}

func (u *productsUsecase) GetAllProducts(ctx context.Context, page, pageSize int) ([]*entity.Product, *dto.PaginationResponse, error) {
	logger.Info().Int("page", page).Int("page_size", pageSize).Msg("Fetching paginated products")

	products, total, err := u.repo.GetAll(ctx, page, pageSize)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch products")
		return nil, nil, err
	}

	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)
	hasNextPage := page < int(totalPages)
	hasPrevPage := page > 1

	nextPage := page + 1
	prevPage := page - 1

	pagination := &dto.PaginationResponse{
		TotalItems:   total,
		TotalPages:   int(totalPages),
		CurrentPage:  page,
		PageSize:     pageSize,
		HasNextPage:  hasNextPage,
		HasPrevPage:  hasPrevPage,
		NextPage:     &nextPage,
		PreviousPage: &prevPage,
	}

	logger.Info().Int("count", len(products)).Int64("total", total).Msg("Products fetched successfully")
	return products, pagination, nil
}
