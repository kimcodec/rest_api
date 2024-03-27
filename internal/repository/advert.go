package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"rest_api/domain"
	"rest_api/lib"
)

type AdvertRepository struct {
	DB *sqlx.DB
}

func NewAdvertRepository(db *sqlx.DB) *AdvertRepository {
	return &AdvertRepository{
		DB: db,
	}
}

func (ar *AdvertRepository) Post(ctx context.Context, adPost domain.AdvertToPost) (domain.AdvertDB, error) {
	conn, err := ar.DB.Connx(ctx)
	if err != nil {
		return domain.AdvertDB{}, err
	}
	defer conn.Close()
	row := conn.QueryRowxContext(
		ctx,
		"INSERT INTO Adverts(user_id, title, text, image_url, price) VALUES ($1, $2, $3, $4, $5) RETURNING *",
		adPost.UserID, adPost.Title, adPost.Text, adPost.ImageURL, adPost.Price)
	if err := row.Err(); err != nil {
		return domain.AdvertDB{}, err
	}

	var ad domain.AdvertDB
	if err := row.StructScan(&ad); err != nil {
		return domain.AdvertDB{}, err
	}

	return ad, nil
}

func (ar *AdvertRepository) List(ctx context.Context, params lib.AdvertQueryParams) ([]domain.AdvertDBWithLogin, error) {
	conn, err := ar.DB.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := fmt.Sprintf("SELECT Adverts.id, user_id, login, title, text, image_url, price, created_at "+
		"FROM Adverts JOIN Users ON user_id = Users.id WHERE Adverts.id >= $1 AND price >= $2 AND price <= $3 "+
		"ORDER BY %s LIMIT $4", func() string {
		switch params.DataSort {
		case domain.FreshDateSort:
			return "created_at DESC"
		case domain.LateDateSort:
			return "created_at"
		case domain.BiggerPriceSort:
			return "price DESC"
		case domain.LessPriceSort:
			return "price"
		default:
			return "created_at"
		}
	}())

	var ads []domain.AdvertDBWithLogin
	if err := conn.SelectContext(ctx, &ads, query, params.Offset,
		params.MinPrice, params.MaxPrice, params.Limit); err != nil {
		return nil, err
	}
	return ads, nil
}
