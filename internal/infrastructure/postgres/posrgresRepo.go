package postgres

import (
	domain "AvitoFlats/internal/domain/user"
	"context"
	"fmt"
	"os/user"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type Repository interface {
	Save(ctx context.Context, u *domain.User) error
	GetByID(ctx context.Context, id string) (*user.User, error)
	GetByEmailOrUsername(ctx context.Context, identifier string) (*user.User, error)
	Update(ctx context.Context, u *user.User) error
	Delete(ctx context.Context, id domain.User) error
}

func (ur *UserRepository) Save(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO user_account(id, email, username, password, time_registration, TimeUpdatedAt) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := ur.db.Exec(ctx, query,
		user.Id, user.Email, user.Username, user.Password, user.TimeRegistration, user.TimeUpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save user in db: %w", err)
	}
	return nil
}

func (ur *UserRepository) GetUserById(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	user := &domain.User{}
	query := `
			SELECT id, email, username, time_registration 
			FROM user_account 
			WHERE id=$1`
	err := ur.db.QueryRow(ctx, query, userID).Scan(&user.Id, &user.Email, &user.Username, &user.TimeRegistration)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found with id: %w", err)
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return user, nil
}

func (ur *UserRepository) GetByEmailOrUsername(ctx context.Context, name string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, email, username, password, time_registration FROM user_account where email=$1 OR username=$1`
	row := ur.db.QueryRow(ctx, query, name)
	err := row.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.TimeRegistration)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found with email/username %w", err)
		}
		return nil, fmt.Errorf("failed to get user by email/username %w", err)
	}
	return user, nil
}
