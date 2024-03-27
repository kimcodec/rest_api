package lib

import (
	"errors"
	"math"
	"net/url"
	"strconv"

	"rest_api/domain"
)

type AdvertQueryParams struct {
	Offset   uint64
	Limit    uint64
	DataSort domain.AdvertSortType
	MinPrice uint64
	MaxPrice uint64
}

func ParseAdvertParams(val url.Values) (AdvertQueryParams, error) {
	var (
		offset   uint64
		limit    uint64
		dataSort domain.AdvertSortType
		minPrice uint64
		maxPrice uint64
	)

	if val.Has("offset") {
		offsetParam := val.Get("offset")
		offsetTemp, err := strconv.Atoi(offsetParam)
		if err != nil {
			return AdvertQueryParams{}, err
		}
		offset = uint64(offsetTemp)
	} else {
		offset = 1
	}

	if val.Has("limit") {
		limitParam := val.Get("limit")
		limitTemp, err := strconv.Atoi(limitParam)
		if err != nil {
			return AdvertQueryParams{}, err
		}
		limit = uint64(limitTemp)
	} else {
		limit = 10
	}

	if val.Has("data_sort") {
		switch val.Get("data_sort") {
		case string(domain.FreshDateSort):
			dataSort = domain.FreshDateSort
		case string(domain.LateDateSort):
			dataSort = domain.LateDateSort
		case string(domain.BiggerPriceSort):
			dataSort = domain.BiggerPriceSort
		case string(domain.LessPriceSort):
			dataSort = domain.LessPriceSort
		default:
			return AdvertQueryParams{}, errors.New("invalid sort param")
		}
	} else {
		dataSort = domain.FreshDateSort
	}

	if val.Has("min_price") {
		minPriceParam := val.Get("min_price")
		minPriceTemp, err := strconv.Atoi(minPriceParam)
		if err != nil {
			return AdvertQueryParams{}, err
		}
		minPrice = uint64(minPriceTemp)
	} else {
		minPrice = 0
	}

	if val.Has("max_price") {
		maxPriceParam := val.Get("max_price")
		maxPriceTemp, err := strconv.Atoi(maxPriceParam)
		if err != nil {
			return AdvertQueryParams{}, err
		}
		offset = uint64(maxPriceTemp)
	} else {
		maxPrice = math.MaxInt
	}

	params := AdvertQueryParams{
		Offset:   offset,
		Limit:    limit,
		DataSort: dataSort,
		MinPrice: minPrice,
		MaxPrice: maxPrice,
	}
	return params, nil
}
