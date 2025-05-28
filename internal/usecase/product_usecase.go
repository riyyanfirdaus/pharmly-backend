package usecase

import (
	"context"
	"pharmly-backend/internal/dto"
	"pharmly-backend/internal/entity"
	"pharmly-backend/internal/logger"
	"pharmly-backend/internal/repository"
	"time"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req *dto.ProductRequest) (*dto.ProductResponse, error)
	GetProductByID(ctx context.Context, id int64) (*entity.Product, error)
	GetAllProducts(ctx context.Context, page, pageSize int) ([]*entity.Product, *dto.PaginationResponse, error)
	UpdateProduct(ctx context.Context, id int64, req *dto.ProductRequest) (*dto.ProductResponse, error)
	DeleteProduct(ctx context.Context, id int64) error
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

func (u *productsUsecase) GetProductByID(ctx context.Context, id int64) (*entity.Product, error) {
	logger.Info().Int64("product_id", id).Msg("Fetching product by ID")

	product, err := u.repo.GetByID(ctx, id)
	if err != nil {
		logger.Error().Err(err).Int64("product_id", id).Msg("Failed to fetch product")
		return nil, err
	}

	logger.Info().Int64("product_id", id).Msg("Product fetched successfully")
	return product, nil
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

func (u *productsUsecase) UpdateProduct(ctx context.Context, id int64, req *dto.ProductRequest) (*dto.ProductResponse, error) {
	logger.Info().Int64("product_id", id).Msg("Starting product update process")

	product, err := u.repo.GetByID(ctx, id)
	if err != nil {
		logger.Error().Err(err).Int64("transaction_id", id).Msg("Failed to fetch transaction")
	}

	product.Name = req.Name
	product.CategoryID = req.CategoryID
	product.GenericName = req.GenericName
	product.Description = &req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	product.Unit = req.Unit
	product.ExpirationDate = req.ExpirationDate
	product.Barcode = req.Barcode
	product.SupplierID = req.SupplierID
	product.MinStock = req.MinStock
	product.IsActive = req.IsActive
	product.UpdatedAt = time.Now()

	if err := u.repo.Update(ctx, product); err != nil {
		logger.Error().Err(err).Int64("product_id", id).Msg("Failed to update product")
		return nil, err
	}

	logger.Info().Int64("product_id", id).Msg("Product updated successfully")
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
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
		DeletedAt:      product.DeletedAt,
	}, nil
}

func (u *productsUsecase) DeleteProduct(ctx context.Context, id int64) error {
	logger.Info().Int64("product_id", id).Msg("Starting product deletion proccess")

	_, err := u.repo.GetByID(ctx, id)
	if err != nil {
		logger.Error().Err(err).Int64("product_id", id).Msg("Failed to fetch product")
		return nil
	}

	if err := u.repo.Delete(ctx, id); err != nil {
		logger.Error().Err(err).Int64("product_id", id).Msg("Failed to delete product")
		return nil
	}

	logger.Info().Int64("product_id", id).Msg("Product deleted successfully")
	return nil
}
