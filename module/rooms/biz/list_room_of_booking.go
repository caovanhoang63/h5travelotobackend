package roombiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	roommodel "h5travelotobackend/module/rooms/model"
)

type ListBookingDetailStore interface {
	ListRoomOfBooking(ctx context.Context, bookingId int) ([]int, error)
}

type listRoomOfBookingBiz struct {
	detailStore  ListBookingDetailStore
	bookingStore FindBookingStore
	roomStore    ListRoomStore
}

func NewListRoomOfBookingBiz(detailStore ListBookingDetailStore, bookingStore FindBookingStore, roomStore ListRoomStore) *listRoomOfBookingBiz {
	return &listRoomOfBookingBiz{detailStore: detailStore, bookingStore: bookingStore, roomStore: roomStore}
}

func (biz *listRoomOfBookingBiz) ListRoomOfBooking(ctx context.Context, bookingId int) ([]roommodel.Room, error) {
	booking, err := biz.bookingStore.FindDTOWithCondition(ctx, map[string]interface{}{"id": bookingId})

	if err != nil {
		return nil, common.ErrEntityNotFound("Booking", err)
	}

	if booking.Status == 0 {
		return nil, common.ErrEntityDeleted("Booking", nil)
	}

	roomIds, err := biz.detailStore.ListRoomOfBooking(ctx, bookingId)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	rooms, err := biz.roomStore.ListRoomInIds(ctx, nil, roomIds)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return rooms, nil
}
