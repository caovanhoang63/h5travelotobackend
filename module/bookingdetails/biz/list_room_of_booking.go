package bookingdetailbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	bookingdetailmodel "h5travelotobackend/module/bookingdetails/model"
)

type ListRoomStore interface {
	ListRoomsOfBooking(ctx context.Context, bookingId int) ([]bookingdetailmodel.Room, error)
}

type listRoomOfBookingBiz struct {
	store ListRoomStore
}

func NewListRoomOfBookingBiz(store ListRoomStore) *listRoomOfBookingBiz {
	return &listRoomOfBookingBiz{
		store: store,
	}
}

func (biz *listRoomOfBookingBiz) ListRoomOfBooking(ctx context.Context, bookingId int) ([]bookingdetailmodel.Room, error) {
	rooms, err := biz.store.ListRoomsOfBooking(ctx, bookingId)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return rooms, nil
}
