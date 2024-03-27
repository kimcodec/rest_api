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
	if err := row.Err(); err != nil {
		return domain.UserDB{}, err
	}

	var user domain.UserDB
	if err := row.StructScan(&user); err != nil {
		return domain.UserDB{}, err
	}
	return user, nil
}

func (ur *UserAuthRepository) GetUserByLogin(ctx context.Context, login string) (domain.UserDB, error) {
	var user []domain.UserDB
	conn, err := ur.db.Connx(ctx)
	if err != nil {
		return domain.UserDB{}, err
	}
	defer conn.Close()
	if err := conn.SelectContext(
		ctx,
		&user,
		"SELECT * FROM USERS WHERE login = $1",
		login); err != nil {
		return domain.UserDB{}, err
	}
	if len(user) == 0 {
		return domain.UserDB{}, errors.New("empty query result")
	}
	return user[0], nil
}
