package domain

import "time"

type AdvertPostRequest struct {
	Title    string `json:"title"`
	Text     string `json:"text"`
	ImageURL string `json:"image_url"`
	Price    string `json:"price"`
}

type AdvertPostResponse struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	ImageURL  string    `json:"image_url"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type AdvertListResponse struct {
	// TODO: подумать над DTO
}

type AdvertToPost struct {
	UserID   uint64
	Title    string
	Text     string
	ImageURL string
	Price    uint64
}

type AdvertDB struct {
	ID        uint64    `db:"id"`
	UserID    uint64    `db:"user_id"`
	Title     string    `db:"title"`
	Text      string    `db:"text"`
	ImageURL  string    `db:"image_url"`
	Price     string    `db:"price"`
	CreatedAt time.Time `db:"created_at"`
}
