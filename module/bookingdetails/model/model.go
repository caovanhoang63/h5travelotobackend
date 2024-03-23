package bookingdetailmodel

import (
	"h5travelotobackend/common"
	"time"
)

const EntityName = "BookingDetail"

type BookingDetail struct {
	BookingId     int        `json:"-" gorm:"column:booking_id"`
	BookingFakeId common.UID `json:"booking_id" gorm:"-"`
	RoomId        int        `json:"-" gorm:"column:room_id"`
	RoomFakeId    common.UID `json:"room_id" gorm:"-"`
	Status        int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt     *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (BookingDetail) TableName() string {
	return "booking_details"
}

func (b *BookingDetail) Mask(isAdmin bool) {
	b.BookingFakeId = common.NewUID(uint32(b.BookingId), common.DbTypeBookingDetail, 1)
	b.RoomFakeId = common.NewUID(uint32(b.RoomId), common.DbTypeBookingDetail, 1)
}

func (b *BookingDetail) UnMask() {
	b.BookingId = int(b.BookingFakeId.GetLocalID())
	b.RoomId = int(b.RoomFakeId.GetLocalID())
}

type BookingDetailCreate BookingDetail

type BookingDetailUpdate struct {
	RoomFakeId common.UID `json:"room_id" gorm:"-"`
	RoomId     int        `json:"-" gorm:"column:room_id"`
}
