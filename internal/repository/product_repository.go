package repository

import (
	"context"
	"pharmly-backend/internal/constant"
	"pharmly-backend/internal/entity"
	"pharmly-backend/internal/logger"
	"time"

	"github.com/jackc/pgx/v5"
)

type ProductRepository interface {
	Create(ctx context.Context, product *entity.Product) error
	GetAll(ctx context.Context, page, pageSize int) ([]*entity.Product, int64, error)
}

type productRepository struct {
	db *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, product *entity.Product) error {
	logger.Info().Str("product", product.Name).Msg("Creating new product")

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to begin transaction")
		return err
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, constant.QCreateProduct, product.Name, product.CategoryID, product.GenericName, product.Description, product.Price, product.Stock, product.Unit, product.ExpirationDate, product.Barcode, product.SupplierID, product.MinStock, product.IsActive, time.Now(), time.Now()).Scan(&product.ID)
	if err != nil {
		logger.Error().Err(err).Int64("product_id", product.ID).Msg("Failed to create product")
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error().Err(err).Msg("Failed to commit transaction")
		return err
	}

	logger.Info().Str("product", product.Name).Msg("Product created successfully")
	return nil
}

func (r *productRepository) GetAll(ctx context.Context, page, pageSize int) ([]*entity.Product, int64, error) {
	logger.Info().Int("page", page).Int("page_size", pageSize).Msg("Fetching paginated products")

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to begin transaction")
		return nil, 0, err
	}
	defer tx.Rollback(ctx)

	var total int64
	err = tx.QueryRow(ctx, constant.QCountProductQuery).Scan(&total)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get total products count")
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	rows, err := tx.Query(ctx, constant.QGetAllProducts, pageSize, offset)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch products")
		return nil, 0, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		product := &entity.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.CategoryID,
			&product.GenericName,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.Unit,
			&product.ExpirationDate,
			&product.Barcode,
			&product.SupplierID,
			&product.MinStock,
			&product.IsActive,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.DeletedAt,
		)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to scan products row")
			return nil, 0, err
		}
		products = append(products, product)
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error().Err(err).Msg("Failed to commit transaction")
		return nil, 0, err
	}

	logger.Info().Int("count", len(products)).Int64("total", total).Msg("products fetch successfully")
	return products, total, nil

}
