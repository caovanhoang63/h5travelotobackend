package bookingdetailmodel

import (
	"fmt"
	"h5travelotobackend/common"
)

func (b *BookingDetailRequest) CheckInvalidBooking(booking *common.DTOBooking) error {
	if booking == nil {
		return ErrBookingNotFound
	}

	if booking.Status == common.StatusDeleted {
		return ErrBookingNotFound
	}

	if len(b.RoomIds) != booking.RoomQuantity {
		return ErrRoomQuantityExceeded
	}

	return nil
}

func (b *BookingDetailRequest) CheckInvalidRoom(rooms []common.DTORoom, RoomTypeId int) error {
	if len(rooms) != len(b.RoomIds) {
		return ErrRoomNotFound
	}

	for i := range rooms {
		if rooms[i].RoomTypeID != RoomTypeId {
			fmt.Println("Room_id", rooms[i].Id)
			fmt.Println("rooms[i].RoomTypeID: ", rooms[i].RoomTypeID)

			return ErrRoomNotFound

		}

	}

	return nil
}
