package repository

import (
	"context"
	"pharmly-backend/internal/constant"
	"pharmly-backend/internal/entity"
	"pharmly-backend/internal/logger"

	"github.com/jackc/pgx/v5"
)

type CategoryRepository interface {
	GetAll(ctx context.Context, page, pageSize int) ([]*entity.Category, int64, error)
}

type categoryRepository struct {
	db *pgx.Conn
}

func NewCategoryRepository(db *pgx.Conn) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAll(ctx context.Context, page, pageSize int) ([]*entity.Category, int64, error) {
	logger.Info().Int("page", page).Int("page_size", pageSize).Msg("Fetching paginated categories")

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to begin transaction")
		return nil, 0, err
	}
	defer tx.Rollback(ctx)

	var total int64
	err = tx.QueryRow(ctx, constant.QCountCategoryQuery).Scan(&total)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get total users count")
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	rows, err := tx.Query(ctx, constant.QGetAllCategories, pageSize, offset)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch users")
		return nil, 0, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		category := &entity.Category{}
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.ParentCategoryID,
			&category.CreatedAt,
			&category.UpdatedAt,
			&category.DeletedAt,
		)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to scan users row")
			return nil, 0, err
		}
		categories = append(categories, category)
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error().Err(err).Msg("Failed to commit transaction")
		return nil, 0, err
	}

	logger.Info().Int("count", len(categories)).Int64("total", total).Msg("categories fetch successfully")
	return categories, total, nil
}
