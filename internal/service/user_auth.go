package service

import (
	"context"
	"rest_api/domain"
)

type UserAuthRepository interface {
	Register(context.Context, domain.UserRegisterRequest) (domain.UserDB, error)
	GetUserByLogin(context.Context, string) (domain.UserDB, error)
}

type UserAuthService struct {
	ur UserAuthRepository
}

func NewUserService(ur UserAuthRepository) *UserAuthService {
	return &UserAuthService{
		ur: ur,
	}
}

func (uc *UserAuthService) GetUserByLogin(ctx context.Context, login string) (domain.User, error) {
	usr, err := uc.ur.GetUserByLogin(ctx, login)
	if err != nil {
		return domain.User{}, err
	}
	user := domain.User{
		ID:       usr.ID,
		Login:    usr.Login,
		Password: usr.Password,
	}
	return user, nil
}

func (uc *UserAuthService) Register(ctx context.Context, req domain.UserRegisterRequest) (domain.UserRegisterResponse, error) {
	userDB, err := uc.ur.Register(ctx, req)
	if err != nil {
		return domain.UserRegisterResponse{}, err
	}
	userResp := domain.UserRegisterResponse{
		ID:       userDB.ID,
		Login:    userDB.Login,
		Password: userDB.Password,
	}
	return userResp, nil
}
