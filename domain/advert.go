package domain

import (
	"time"
)

type AdvertPostRequest struct {
	Title    string `json:"title" validate:"required,min=4,max=50"`
	Text     string `json:"text" validate:"required,min=4,max=500"`
	ImageURL string `json:"image_url" validate:"required,url"`
	Price    string `json:"price" validate:"required,number"`
}

type AdvertPostResponse struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	ImageURL  string    `json:"image_url"`
	Price     uint64    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type AdvertSortType string

const (
	FreshDateSort   AdvertSortType = "fresh_date"
	LateDateSort    AdvertSortType = "late_date"
	BiggerPriceSort AdvertSortType = "bigger_price"
	LessPriceSort   AdvertSortType = "less_price"
)

type AdvertListResponse struct {
	AuthorLogin     string `json:"author_login"`
	Title           string `json:"title"`
	Text            string `json:"text"`
	ImageURL        string `json:"image_url"`
	Price           uint64 `json:"price"`
	PublishedByUser bool   `json:"published_by_user"`
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
	Price     uint64    `db:"price"`
	CreatedAt time.Time `db:"created_at"`
}

type AdvertDBWithLogin struct {
	ID        uint64    `db:"id"`
	UserID    uint64    `db:"user_id"`
	UserLogin string    `db:"login"`
	Title     string    `db:"title"`
	Text      string    `db:"text"`
	ImageURL  string    `db:"image_url"`
	Price     uint64    `db:"price"`
	CreatedAt time.Time `db:"created_at"`
}
