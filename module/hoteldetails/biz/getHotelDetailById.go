package hoteldetailbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hoteldetailmodel "h5travelotobackend/module/hoteldetails/model"
)

type GetHotelDetailByIdStore interface {
	FindWithCondition(ctx context.Context, condition map[string]interface{}) (*hoteldetailmodel.HotelDetail, error)
}

type getHotelDetailByIdBiz struct {
	store GetHotelDetailByIdStore
}

func NewGetHotelDetailByIdBiz(store GetHotelDetailByIdStore) *getHotelDetailByIdBiz {
	return &getHotelDetailByIdBiz{store: store}
}

func (biz *getHotelDetailByIdBiz) GetHotelDetailById(ctx context.Context, hotelId int) (*hoteldetailmodel.HotelDetail, error) {
	data, err := biz.store.FindWithCondition(ctx, map[string]interface{}{"hotel_id": hotelId})
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrEntityNotFound(hoteldetailmodel.EntityName, err)
		} else {
			return nil, common.ErrInternal(err)
		}
	}
	return data, nil
}
