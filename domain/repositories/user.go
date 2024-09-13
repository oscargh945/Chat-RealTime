package repositories

import (
	"context"
	"github.com/oscargh945/go-Chat/domain/entities"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
}

type UserService interface {
	CreateUserService(ctx context.Context, req *entities.CreateUserReq) (*entities.CreateUserResp, error)
}
