package hotelfacilitiesbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelfacilitymodel "h5travelotobackend/module/hotelfacilities/model"
)

type ListFaciOfHotelStore interface {
	ListFacilitiesOfHotel(ctx context.Context, hotelId int) ([]hotelfacilitymodel.HotelFacility, error)
}

type listFaciOfHotelBiz struct {
	store ListFaciOfHotelStore
}

func NewListFaciOfHotelBiz(store ListFaciOfHotelStore) *listFaciOfHotelBiz {
	return &listFaciOfHotelBiz{store: store}
}

func (biz *listFaciOfHotelBiz) ListFacilitiesOfHotel(ctx context.Context, hotelId int) ([]hotelfacilitymodel.HotelFacility, error) {
	data, err := biz.store.ListFacilitiesOfHotel(ctx, hotelId)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return data, nil
}
