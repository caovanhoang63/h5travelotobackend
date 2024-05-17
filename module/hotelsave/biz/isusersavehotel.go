package htsavebiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htsavemodel "h5travelotobackend/module/hotelsave/model"
)

type FindSavedHotelStore interface {
	FindSavedHotel(ctx context.Context, conditions map[string]interface{}) (*htsavemodel.HotelSave, error)
}

type FindSavedHotelBiz struct {
	store FindSavedHotelStore
}

func NewFindSavedHotelBiz(store FindSavedHotelStore) *FindSavedHotelBiz {
	return &FindSavedHotelBiz{store: store}
}

func (biz *FindSavedHotelBiz) IsHotelSaved(ctx context.Context, requester common.Requester, hotelId int) bool {
	conditions := map[string]interface{}{"user_id": requester.GetUserId(), "hotel_id": hotelId}
	_, err := biz.store.FindSavedHotel(ctx, conditions)
	if err != nil {
		return false
	}
	return true
}
