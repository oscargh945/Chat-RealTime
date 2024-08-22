package repositories

import (
	"context"
	"github.com/oscargh945/go-Chat/domain/entities"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	ListUsers(ctx context.Context) (*[]entities.User, error)
}

type UserService interface {
	CreateUserService(ctx context.Context, req *entities.CreateUserReq) (*entities.CreateUserResp, error)
	ListUsersService(ctx context.Context) (*[]entities.User, error)
}
