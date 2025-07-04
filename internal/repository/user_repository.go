package repository

import (
	"context"
	"errors"
	"pharmly-backend/internal/constant"
	"pharmly-backend/internal/entity"
	"pharmly-backend/internal/logger"
	"time"

	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetAll(ctx context.Context, page, pageSize int) ([]*entity.User, int64, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByID(ctx context.Context, id int64) (*entity.User, error)
}

type userRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	logger.Info().Str("email", user.Email).Msg("Creating new user")

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to begin transaction")
		return err
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, constant.QCreateUser, user.Username, user.FullName, user.Email, user.Password, user.Role, time.Now(), time.Now()).Scan(&user.ID)
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)" {
			return errors.New("email already exists")
		}
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)" {
			return errors.New("username already exists")
		}
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error().Err(err).Msg("Failed to commit transaction")
		return err
	}

	logger.Info().Str("email", user.Email).Int64("user_id", user.ID).Msg("User created successfully")
	return nil
}

func (r *userRepository) GetAll(ctx context.Context, page, pageSize int) ([]*entity.User, int64, error) {
	logger.Info().Int("page", page).Int("page_size", pageSize).Msg("Fetching paginated users")

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to begin transaction")
		return nil, 0, err
	}
	defer tx.Rollback(ctx)

	var total int64
	err = tx.QueryRow(ctx, constant.QCountUserQuery).Scan(&total)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get total users count")
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	rows, err := tx.Query(ctx, constant.QGetAllUsers, pageSize, offset)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch users")
		return nil, 0, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		user := &entity.User{}
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.FullName,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.Status,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to scan users row")
			return nil, 0, err
		}
		users = append(users, user)
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error().Err(err).Msg("Failed to commit transaction")
		return nil, 0, err
	}

	logger.Info().Int("count", len(users)).Int64("total", total).Msg("Users fetch successfully")
	return users, total, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	logger.Info().Str("email", email).Msg("Fetching user by email")

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to begin transaction")
		return nil, err
	}
	defer tx.Rollback(ctx)

	user := &entity.User{}
	err = tx.QueryRow(ctx, constant.QGetByEmail, email).Scan(
		&user.ID,
		&user.Username,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err == pgx.ErrNoRows {
		logger.Error().Str("email", email).Msg("User not found")
		return nil, nil
	}

	if err != nil {
		logger.Error().Err(err).Str("email", email).Msg("Failed to fetch user")
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error().Err(err).Msg("Failed to commit transaction")
		return nil, err
	}

	logger.Info().Str("email", email).Int64("user_id", user.ID).Msg("User fetched successfully")
	return user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	logger.Info().Int64("user_id", id).Msg("Fetching user by ID")

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to begin transaction")
		return nil, err
	}
	defer tx.Rollback(ctx)

	user := &entity.User{}
	err = tx.QueryRow(ctx, constant.QGetByID, id).Scan(
		&user.ID,
		&user.Username,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.Password,
		&user.Role,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err == pgx.ErrNoRows {
		logger.Error().Int64("user_id", id).Msg("User not found")
		return nil, nil
	}

	if err != nil {
		logger.Error().Err(err).Int64("user_id", id).Msg("Failed to fetch user")
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error().Err(err).Msg("Failed to commit transaction")
		return nil, err
	}

	logger.Info().Int64("user_id", id).Msg("User fetched successfully")
	return user, nil
}
