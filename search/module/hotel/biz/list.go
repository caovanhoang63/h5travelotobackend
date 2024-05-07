package hotelsearchbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/search/module/hotel/model"
)

type ListHotelStore interface {
	ListHotel(ctx context.Context, filter *hotelmodel.Filter, paging *common.Paging) ([]hotelmodel.Hotel, error)
}

type listHotelBiz struct {
	store ListHotelStore
}

func NewListHotelBiz(store ListHotelStore) *listHotelBiz {
	return &listHotelBiz{store: store}
}

func (biz *listHotelBiz) ListHotelWithFilter(ctx context.Context,
	filter *hotelmodel.Filter,
	paging *common.Paging) ([]hotelmodel.Hotel, error) {
	err := filter.Validate()
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	result, err := biz.store.ListHotel(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return result, nil
}
