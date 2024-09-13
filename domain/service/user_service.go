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
		UserName: req.UserName,
		Email:    req.Email,
		Password: hashpassword,
	}

	r, err := s.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	result := &entities.CreateUserResp{
		ID:       r.ID,
		UserName: r.UserName,
		Email:    r.Email,
	}
	return result, nil
}

func (s *UserService) Login(ctx context.Context, entity *entities.LoginRequest) (*entities.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	user, err := s.GetUserByEmail(ctx, entity.Email)
	if err != nil {
		return &entities.LoginResponse{}, err
	}

	err = utils.CheckPassword(entity.Password, user.Password)
	if err != nil {
		return &entities.LoginResponse{}, err
	}

	tokens, err := GenerateTokens(*user)
	if err != nil {
		return &entities.LoginResponse{}, err
	}

	return &tokens, nil
}

func (s *UserService) RefreshTokenUserService(refreshToken string) (*entities.LoginResponse, error) {
	tokens, err := RefreshToken(refreshToken)
	if err != nil {
		return &entities.LoginResponse{}, err
	}

	_, err = ValidateToken(tokens.AccessToken)
	if err != nil {
		return &entities.LoginResponse{}, err
	}
	return tokens, nil
}
