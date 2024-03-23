package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"rest_api/domain"
)

type UserRepository struct {
	Conn *sqlx.Conn
}

func NewUserRepository(conn *sqlx.Conn) *UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

func (ur *UserRepository) Register(ctx context.Context, req domain.UserRegisterRequest) error {
	if _, err := ur.Conn.ExecContext(ctx,
		"INSERT INTO USERS VALUES($1, $2)",
		req.Login, req.Password); err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserByLogin(ctx context.Context, login string) (domain.User, error) {
	var user domain.User
	if err := ur.Conn.SelectContext(ctx, &user,
		"SELECT login, password FROM USERS WHERE login = $1",
		login); err != nil {
		return user, err
	}
	return user, nil
}
