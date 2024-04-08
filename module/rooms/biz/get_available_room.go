package roombiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	roommodel "h5travelotobackend/module/rooms/model"
	"time"
)

type ListRoomBookedRepo interface {
	GetRoomIdsBooked(
		ctx context.Context,
		startDate *time.Time,
		endDate *time.Time,
		condition map[string]interface{}) ([]int, error)
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

func (biz *listAvailableRoomBiz) ListAvailableRoomForBooking(ctx context.Context, bookingId int) ([]roommodel.Room, error) {
	booking, err := biz.bookingStore.FindDTOWithCondition(ctx, map[string]interface{}{"id": bookingId})

	if err != nil {
		return nil, common.ErrEntityNotFound("Booking", err)
	}

	roomIds, err := biz.repo.GetRoomIdsBooked(ctx, booking.StartDate, booking.EndDate, map[string]interface{}{"bookings.room_type_id": booking.RoomTypeId})
	if err != nil {
		return nil, common.ErrCannotListEntity("Room", err)
	}

	condition := map[string]interface{}{
		"room_type_id": booking.RoomTypeId,
		"condition":    "available",
	}

	rooms, err := biz.roomStore.ListRoomsNotInIds(ctx, condition, roomIds)
	if err != nil {
		return nil, common.ErrCannotListEntity("Room", err)
	}

	return rooms, nil
}

func (biz *listAvailableRoomBiz) ListAvailableRoomByDate(
	ctx context.Context,
	startDate *time.Time,
	endDate *time.Time,
	hotelId int,
) ([]roommodel.Room, error) {

	roomIds, err := biz.repo.GetRoomIdsBooked(ctx,
		startDate,
		endDate,
		map[string]interface{}{
			"bookings.hotel_id": hotelId,
		})
	if err != nil {
		return nil, common.ErrCannotListEntity("Room", err)
	}

	condition := map[string]interface{}{"condition": "available"}

	rooms, err := biz.roomStore.ListRoomsNotInIds(ctx, condition, roomIds)
	if err != nil {
		return nil, common.ErrCannotListEntity("Room", err)
	}

	return rooms, nil
}
