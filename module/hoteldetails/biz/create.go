package hoteldetailbiz

import (
	"context"
	"h5travelotobackend/common"
	hoteldetailmodel "h5travelotobackend/module/hoteldetails/model"
)

type CreateHotelDetailStore interface {
	Create(ctx context.Context, data *hoteldetailmodel.HotelDetail) error
}

type createHotelDetailBiz struct {
	store CreateHotelDetailStore
}

func NewCreateHotelDetailBiz(store CreateHotelDetailStore) *createHotelDetailBiz {
	return &createHotelDetailBiz{store: store}
}

func (biz *createHotelDetailBiz) CreateHotelDetail(ctx context.Context, data *hoteldetailmodel.HotelDetail) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(hoteldetailmodel.EntityName, err)
	}
	return nil
}
