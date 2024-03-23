package bookingdetailmodel

import "h5travelotobackend/common"

type Filter struct {
	BookingId     int        `json:"-" gorm:"column:booking_id"`
	BookingFakeId common.UID `json:"booking_id" gorm:"-"`
	RoomId        int        `json:"-" gorm:"column:room_id"`
	RoomFakeId    common.UID `json:"room_id" gorm:"-"`
}

func (f *Filter) Mask(isAdmin bool) {
	f.BookingFakeId = common.NewUID(uint32(f.BookingId), common.DbTypeBookingDetail, 1)
	f.RoomFakeId = common.NewUID(uint32(f.RoomId), common.DbTypeBookingDetail, 1)
}

func (f *Filter) UnMask() {
	f.BookingId = int(f.BookingFakeId.GetLocalID())
	f.RoomId = int(f.RoomFakeId.GetLocalID())
}
