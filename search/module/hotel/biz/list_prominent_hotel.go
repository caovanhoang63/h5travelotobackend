package hotelsearchbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/search/module/hotel/model"
)

type ListProminentHotelStore interface {
	ListRandomHotels(ctx context.Context, limit int) ([]hotelmodel.Hotel, error)
}

type GetMinPrice interface {
	GetMinPriceByHotelId(ctx context.Context, hotelId int) (float64, error)
}

type listProminentHotelBiz struct {
	store   ListProminentHotelStore
	rStore  GetMinPrice
	rtStore GetRatingStore
}

func NewListProminentHotelBiz(store ListProminentHotelStore, rStore GetMinPrice, rtStore GetRatingStore) *listProminentHotelBiz {
	return &listProminentHotelBiz{store: store, rStore: rStore, rtStore: rtStore}
}

func (biz *listProminentHotelBiz) ListProminentHotel(ctx context.Context, limit int) ([]hotelmodel.Hotel, error) {
	result, err := biz.store.ListRandomHotels(ctx, limit)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	for i := len(result) - 1; i >= 0; i-- {
		p, err := biz.rStore.GetMinPriceByHotelId(ctx, result[i].Id)
		if err != nil {
			result = append(result[:i], result[i+1:]...)
			continue
		}
		result[i].DisplayPrice = &p
		result[i].Rating, result[i].TotalRating, _ = biz.rtStore.GetHotelRating(ctx, result[i].Id)

	}

	return result, nil
}
