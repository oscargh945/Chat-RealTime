package service

import (
	"context"
	"github.com/oscargh945/go-Chat/domain/entities"
	"github.com/oscargh945/go-Chat/domain/repositories"
	"github.com/oscargh945/go-Chat/utils"
	"time"
)

type UserService struct {
	repositories.UserRepository
	timeout time.Duration
}

func NewUserService(repository repositories.UserRepository) *UserService {
	return &UserService{
		repository,
		2 * time.Second,
	}
}

func (s *UserService) CreateUserService(ctx context.Context, req *entities.CreateUserReq) (*entities.CreateUserResp, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	hashpassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashpassword,
	}

	r, err := s.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	result := &entities.CreateUserResp{
		ID:       r.ID,
		Username: r.Username,
		Email:    r.Email,
	}
	return result, nil
}
