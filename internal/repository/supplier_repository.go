package repository

import (
	"context"
	"pharmly-backend/internal/constant"
	"pharmly-backend/internal/entity"
	"pharmly-backend/internal/logger"

	"github.com/jackc/pgx/v5"
)

type SupplierRepository interface {
	GetAll(ctx context.Context, page, pageSize int) ([]*entity.Supplier, int64, error)
}

type supplierRepository struct {
	db *pgx.Conn
}

func NewSupplierRepository(db *pgx.Conn) SupplierRepository {
	return &supplierRepository{db: db}
}

func (r *supplierRepository) GetAll(ctx context.Context, page, pageSize int) ([]*entity.Supplier, int64, error) {
	logger.Info().Int("page", page).Int("page_size", pageSize).Msg("Fetching paginated products")

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to begin transaction")
		return nil, 0, err
	}
	defer tx.Rollback(ctx)

	var total int64
	err = tx.QueryRow(ctx, constant.QCountSupplierQuery).Scan(&total)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get total suppliers count")
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	rows, err := tx.Query(ctx, constant.QGetAllSuppliers, pageSize, offset)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch suppliers")
		return nil, 0, err
	}
	defer rows.Close()

	var suppliers []*entity.Supplier
	for rows.Next() {
		supplier := &entity.Supplier{}
		err := rows.Scan(
			&supplier.ID,
			&supplier.Name,
			&supplier.ContactPerson,
			&supplier.Phone,
			&supplier.Address,
			&supplier.Email,
			&supplier.CreatedAt,
			&supplier.UpdatedAt,
			&supplier.DeletedAt,
		)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to scan suppliers row")
			return nil, 0, err
		}
		suppliers = append(suppliers, supplier)
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error().Err(err).Msg("Failed to commit transaction")
		return nil, 0, err
	}

	logger.Info().Int("count", len(suppliers)).Int64("total", total).Msg("Suppliers fetch successfully")
	return suppliers, total, nil
}
