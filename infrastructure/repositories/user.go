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
	query := "INSERT INTO users (user_name, email, password) VALUES ($1, $2, $3) RETURNING id, user_name, email"
	row := r.Pool.QueryRow(ctx, query, user.UserName, user.Email, user.Password)

	var id uuid.UUID
	var userName, email string
	if err := row.Scan(&id, &userName, &email); err != nil {
		return nil, err
	}

	user.ID = id
	user.UserName = userName
	user.Email = email

	return user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := "SELECT id, user_name, password FROM users WHERE email = $1"
	row := r.Pool.QueryRow(ctx, query, email)

	var id uuid.UUID
	var userName, password string
	if err := row.Scan(&id, &userName, &password); err != nil {
		return nil, err
	}
	user := &entities.User{
		ID:       id,
		UserName: userName,
		Password: password,
	}
	return user, nil
}
