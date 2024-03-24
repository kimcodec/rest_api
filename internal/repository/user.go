package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"rest_api/domain"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Register(ctx context.Context, req domain.UserRegisterRequest) error {
	conn, err := ur.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	if _, err := conn.ExecContext(ctx,
		"INSERT INTO USERS VALUES($1, $2)",
		req.Login, req.Password); err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserByLogin(ctx context.Context, login string) (domain.User, error) {
	var user []domain.User
	conn, err := ur.db.Connx(ctx)
	if err != nil {
		return domain.User{}, err
	}
	defer conn.Close()
	if err := conn.SelectContext(ctx, &user,
		"SELECT login, password FROM USERS WHERE login = $1",
		login); err != nil {
		return domain.User{}, err
	}
	return user[0], nil
}
