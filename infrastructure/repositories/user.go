package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oscargh945/go-Chat/domain/entities"
)

type UserRepository struct {
	Pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool, ctx context.Context) *UserRepository {
	return &UserRepository{
		Pool: pool,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email"
	row := r.Pool.QueryRow(ctx, query, user.Username, user.Email, user.Password)

	var id uuid.UUID
	var username, email string
	if err := row.Scan(&id, &username, &email); err != nil {
		return nil, err
	}

	user.ID = id
	user.Username = username
	user.Email = email

	return user, nil
}
