package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"rest_api/domain"
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
	if row.Err() != nil {
		return domain.AdvertDB{}, err
	}

	var ad domain.AdvertDB
	if err := row.Scan(&ad); err != nil {
		return domain.AdvertDB{}, err
	}

	return ad, nil
}
