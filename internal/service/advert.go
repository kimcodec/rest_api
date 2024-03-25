package service

import (
	"context"
	"rest_api/domain"
)

type AdvertRepository interface {
	Post(context.Context, domain.AdvertToPost) (domain.AdvertDB, error)
}

type AdvertService struct {
	ad AdvertRepository
}

func NewAdvertService(ad AdvertRepository) *AdvertService {
	return &AdvertService{
		ad: ad,
	}
}

func (as *AdvertService) Post(ctx context.Context, ad domain.AdvertToPost) (domain.AdvertPostResponse, error) {
	respDB, err := as.ad.Post(ctx, ad)
	if err != nil {
		return domain.AdvertPostResponse{}, err
	}

	resp := domain.AdvertPostResponse{
		ID:        respDB.ID,
		UserID:    respDB.UserID,
		Title:     respDB.Title,
		Text:      respDB.Text,
		ImageURL:  respDB.ImageURL,
		Price:     respDB.Price,
		CreatedAt: respDB.CreatedAt,
	}
	return resp, nil
}
