package repository

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"rest_api/domain"
)

type UserAuthRepository struct {
	db *sqlx.DB
}

func NewUserAuthRepository(db *sqlx.DB) *UserAuthRepository {
	return &UserAuthRepository{
		db: db,
	}
}

func (ur *UserAuthRepository) Register(ctx context.Context, req domain.UserRegisterRequest) (domain.UserDB, error) {
	conn, err := ur.db.Connx(ctx)
	if err != nil {
		return domain.UserDB{}, err
	}
	defer conn.Close()
	if _, err := conn.ExecContext(ctx,
		"INSERT INTO Users(login, password) VALUES($1, $2)",
		req.Login, req.Password); err != nil {
		return domain.UserDB{}, err
	}
	var user []domain.UserDB
	if err := conn.SelectContext(ctx, &user, "SELECT * FROM Users WHERE login = $1", req.Login); err != nil {
		return domain.UserDB{}, err
	}
	return user[0], nil
}

func (ur *UserAuthRepository) GetUserByLogin(ctx context.Context, login string) (domain.UserGetByLogin, error) {
	var user []domain.UserGetByLogin
	conn, err := ur.db.Connx(ctx)
	if err != nil {
		return domain.UserGetByLogin{}, err
	}
	defer conn.Close()
	if err := conn.SelectContext(ctx, &user,
		"SELECT login, password FROM USERS WHERE login = $1",
		login); err != nil {
		return domain.UserGetByLogin{}, err
	}
	if len(user) == 0 {
		return domain.UserGetByLogin{}, errors.New("empty query result")
	}
	return user[0], nil
}
