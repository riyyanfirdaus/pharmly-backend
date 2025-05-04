package usecase

import (
	"context"
	"pharmly-backend/internal/dto"
	"pharmly-backend/internal/entity"
	"pharmly-backend/internal/logger"
	"pharmly-backend/internal/repository"
)

type UserUsecase interface {
	GetAllUsers(ctx context.Context, page, pageSize int) ([]*entity.User, *dto.PaginationResponse, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) GetAllUsers(ctx context.Context, page, pageSize int) ([]*entity.User, *dto.PaginationResponse, error) {
	logger.Info().Int("page", page).Int("page_size", pageSize).Msg("Fetching paginated users")

	users, total, err := u.repo.GetAll(ctx, page, pageSize)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch users")
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

	logger.Info().Int("count", len(users)).Int64("total", total).Msg("Users fetched successfully")
	return users, pagination, nil
}
