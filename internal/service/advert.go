package service

import (
	"context"

	"rest_api/domain"
	"rest_api/lib"
)

type AdvertRepository interface {
	Post(context.Context, domain.AdvertToPost) (domain.AdvertDB, error)
	List(context.Context, lib.AdvertQueryParams) ([]domain.AdvertDBWithLogin, error)
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

func (as *AdvertService) List(ctx context.Context, params lib.AdvertQueryParams, id uint64) ([]domain.AdvertListResponse, error) {
	adDB, err := as.ad.List(ctx, params)
	if err != nil {
		return nil, err
	}

	adResp := make([]domain.AdvertListResponse, 0, len(adDB))
	for _, v := range adDB {
		tmp := domain.AdvertListResponse{
			AuthorLogin: v.UserLogin,
			Title:       v.Title,
			Text:        v.Text,
			ImageURL:    v.ImageURL,
			Price:       v.Price,
			PublishedByUser: func() bool {
				if v.UserID == id {
					return true
				}
				return false
			}(),
		}
		adResp = append(adResp, tmp)
	}
	return adResp, err
}
