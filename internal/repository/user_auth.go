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

	row := conn.QueryRowxContext(
		ctx,
		"INSERT INTO Users(login, password) VALUES($1, $2) RETURNING *",
		req.Login, req.Password)
	if row.Err() != nil {
		return domain.UserDB{}, err
	}

	var user domain.UserDB
	if err := row.Scan(&user); err != nil {
		return domain.UserDB{}, err
	}
	return user, nil
}

func (ur *UserAuthRepository) GetUserByLogin(ctx context.Context, login string) (domain.UserGetByLogin, error) {
	var user []domain.UserGetByLogin
	conn, err := ur.db.Connx(ctx)
	if err != nil {
		return domain.UserGetByLogin{}, err
	}
	defer conn.Close()
	if err := conn.SelectContext(
		ctx,
		&user,
		"SELECT login, password FROM USERS WHERE login = $1",
		login); err != nil {
		return domain.UserGetByLogin{}, err
	}
	if len(user) == 0 {
		return domain.UserGetByLogin{}, errors.New("empty query result")
	}
	return user[0], nil
}
