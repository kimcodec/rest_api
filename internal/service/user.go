package service

import (
	"context"
	"rest_api/domain"
)

type UserRepository interface {
	Register(context.Context, domain.UserRegisterRequest) error
	GetUserByLogin(context.Context, string) (domain.User, error)
}

type UserService struct {
	ur UserRepository
}

func NewUserService(ur UserRepository) *UserService {
	return &UserService{
		ur: ur,
	}
}

func (uc *UserService) GetUserByLogin(ctx context.Context, login string) (domain.User, error) {
	return uc.ur.GetUserByLogin(ctx, login)
}

func (uc *UserService) Register(ctx context.Context, req domain.UserRegisterRequest) error {
	return uc.ur.Register(ctx, req)
}
