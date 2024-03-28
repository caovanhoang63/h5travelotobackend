package roombiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	roommodel "h5travelotobackend/module/rooms/model"
)

type ListRoomBookedRepo interface {
	GetRoomIdsBooked(ctx context.Context, booking *common.DTOBooking) ([]int, error)
}

type FindBookingStore interface {
	FindDTOWithCondition(ctx context.Context,
		condition map[string]interface{}) (*common.DTOBooking, error)
}

type listAvailableRoomBiz struct {
	repo         ListRoomBookedRepo
	roomStore    ListRoomStore
	bookingStore FindBookingStore
}

func NewListAvailableRoom(repo ListRoomBookedRepo, roomStore ListRoomStore, bookingStore FindBookingStore) *listAvailableRoomBiz {
	return &listAvailableRoomBiz{repo: repo, roomStore: roomStore, bookingStore: bookingStore}
}

func (biz *listAvailableRoomBiz) ListAvailableRoom(ctx context.Context, bookingId int) ([]roommodel.Room, error) {
	booking, err := biz.bookingStore.FindDTOWithCondition(ctx, map[string]interface{}{"id": bookingId})

	if err != nil {
		return nil, common.ErrEntityNotFound("Booking", err)
	}

	roomIds, err := biz.repo.GetRoomIdsBooked(ctx, booking)
	if err != nil {
		return nil, common.ErrCannotListEntity("Room", err)
	}

	condition := map[string]interface{}{"room_type_id": booking.RoomTypeId}

	rooms, err := biz.roomStore.ListRoomsNotInIds(ctx, condition, roomIds)
	if err != nil {
		return nil, common.ErrCannotListEntity("Room", err)
	}

	return rooms, nil
}
