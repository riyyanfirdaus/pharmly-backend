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
	GetByID(ctx context.Context, id int64) (*entity.Product, error)
	GetAll(ctx context.Context, page, pageSize int) ([]*entity.Product, int64, error)
	Update(ctx context.Context, product *entity.Product) error
	Delete(ctx context.Context, id int64) error
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

func (r *productRepository) GetByID(ctx context.Context, id int64) (*entity.Product, error) {
	logger.Info().Int64("product_id", id).Msg("Fetching product by ID")

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to begin transaction")
		return nil, err
	}
	defer tx.Rollback(ctx)

	product := &entity.Product{}
	err = tx.QueryRow(ctx, constant.QGetProductByID, id).Scan(
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
		logger.Error().Err(err).Int64("product_id", product.ID).Msg("Failed to fecth product")
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error().Err(err).Msg("Failed to commit transaction")
		return nil, err
	}

	return product, nil
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

func (r *productRepository) Update(ctx context.Context, product *entity.Product) error {
	logger.Info().Int64("product_id", product.ID).Msg("Updating Product")

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to begin transaction")
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, constant.QUpdateProduct,
		product.Name,
		product.CategoryID,
		product.GenericName,
		product.Description,
		product.Price,
		product.Stock,
		product.Unit,
		product.ExpirationDate,
		product.Barcode,
		product.SupplierID,
		product.MinStock,
		product.IsActive,
		time.Now(),
		product.ID,
	)

	if err != nil {
		logger.Error().Err(err).Int64("product_id", product.ID).Msg("Failed to update product")
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error().Err(err).Msg("Failed to commit transaction")
		return err
	}
	logger.Info().Int64("product_id", product.ID).Msg("Product updated successfully")
	return nil
}

func (r *productRepository) Delete(ctx context.Context, id int64) error {
	logger.Info().Int64("product_id", id).Msg("Deleting product")

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to begin transaction")
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, constant.QDeleteProduct, id)
	if err != nil {
		logger.Error().Err(err).Int64("product_id", id).Msg("Failed to delete product")
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error().Err(err).Msg("Failed to commit transaction")
		return err
	}

	logger.Info().Int64("product_id", id).Msg("Product deleted successfully")

	return nil
}
