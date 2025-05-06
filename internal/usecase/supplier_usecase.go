package usecase

import (
	"context"
	"pharmly-backend/internal/dto"
	"pharmly-backend/internal/entity"
	"pharmly-backend/internal/logger"
	"pharmly-backend/internal/repository"
)

type SupplierUsecase interface {
	GetAllSuppliers(ctx context.Context, page, pageSize int) ([]*entity.Supplier, *dto.PaginationResponse, error)
}

type supplierUsecase struct {
	repo repository.SupplierRepository
}

func NewSupplierUsecase(repo repository.SupplierRepository) SupplierUsecase {
	return &supplierUsecase{repo: repo}
}

func (u *supplierUsecase) GetAllSuppliers(ctx context.Context, page, pageSize int) ([]*entity.Supplier, *dto.PaginationResponse, error) {
	logger.Info().Int("page", page).Int("page_size", pageSize).Msg("Fetching paginated suppliers")

	suppliers, total, err := u.repo.GetAll(ctx, page, pageSize)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch suppliers")
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

	logger.Info().Int("count", len(suppliers)).Int64("total", total).Msg("Suppliers fetched successfully")
	return suppliers, pagination, nil
}
