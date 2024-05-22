package hotelsearchbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/search/module/hotel/model"
)

type ListProminentHotelStore interface {
	ListRandomHotels(ctx context.Context, limit int) ([]hotelmodel.Hotel, error)
}
type listProminentHotelBiz struct {
	store ListProminentHotelStore
}

func NewListProminentHotelBiz(store ListProminentHotelStore) *listProminentHotelBiz {
	return &listProminentHotelBiz{store: store}
}

func (biz *listProminentHotelBiz) ListProminentHotel(ctx context.Context, limit int) ([]hotelmodel.Hotel, error) {
	result, err := biz.store.ListRandomHotels(ctx, limit)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return result, nil
}
