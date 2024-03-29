package bookingmodel

import (
	"h5travelotobackend/common"
)

const EntityName = "Booking"

type Booking struct {
	common.SqlModel  `json:",inline"`
	HotelId          int               `json:"-" gorm:"column:hotel_id"`
	HotelFakeId      common.UID        `json:"hotel_id"`
	UserId           int               `json:"-" gorm:"column:user_id"`
	UserFakeId       common.UID        `json:"user_id"`
	RoomTypeId       int               `json:"-" gorm:"column:room_type_id"`
	RoomTypeFakeId   common.UID        `json:"room_type_id" gorm:"-"`
	RoomQuantity     int               `json:"room_quantity" gorm:"column:room_quantity"`
	CustomerQuantity int               `json:"customer_quantity" gorm:"column:customer_quantity"`
	StartDate        *common.CivilTime `json:"start_date" gorm:"column:start_date"`
	EndDate          *common.CivilTime `json:"end_date" gorm:"column:end_date"`
}

func (Booking) TableName() string {
	return "bookings"
}

func (b *Booking) Mask(isAdmin bool) {
	b.GenUID(common.DbTypeBooking)
	b.HotelFakeId = common.NewUID(uint32(b.HotelId), common.DbTypeHotel, 1)
	b.UserFakeId = common.NewUID(uint32(b.UserId), common.DbTypeUser, 1)
	b.RoomTypeFakeId = common.NewUID(uint32(b.RoomTypeId), common.DbTypeRoomType, 1)
}

func (b *Booking) UnMask() {
	b.HotelId = int(b.HotelFakeId.GetLocalID())
	b.UserId = int(b.UserFakeId.GetLocalID())
	b.RoomTypeId = int(b.RoomTypeFakeId.GetLocalID())
}

type BookingCreate struct {
	common.SqlModel  `json:",inline"`
	HotelId          int               `json:"-" gorm:"column:hotel_id"`
	HotelFakeId      common.UID        `json:"hotel_id" gorm:"-"`
	UserId           int               `json:"-" gorm:"column:user_id"`
	RoomTypeId       int               `json:"-" gorm:"column:room_type_id"`
	RoomTypeFakeId   common.UID        `json:"room_type_id" gorm:"-"`
	CustomerQuantity int               `json:"customer_quantity" gorm:"column:customer_quantity"`
	RoomQuantity     int               `json:"room_quantity" gorm:"column:room_quantity"`
	StartDate        *common.CivilTime `json:"start_date" gorm:"column:start_date"`
	EndDate          *common.CivilTime `json:"end_date" gorm:"column:end_date"`
}

func (b *BookingCreate) UnMask() {
	b.HotelId = int(b.HotelFakeId.GetLocalID())
	b.RoomTypeId = int(b.RoomTypeFakeId.GetLocalID())
}

func (b *BookingCreate) Mask(isAdmin bool) {
	b.GenUID(common.DbTypeBooking)
	b.HotelFakeId = common.NewUID(uint32(b.HotelId), common.DbTypeHotel, 1)
	b.RoomTypeId = int(b.RoomTypeFakeId.GetLocalID())
}

func (BookingCreate) TableName() string {
	return Booking{}.TableName()
}

type BookingUpdate struct {
	CustomerQuantity int               `json:"customer_quantity" gorm:"column:customer_quantity"`
	StartDate        *common.CivilTime `json:"start_date" gorm:"column:start_date"`
	EndDate          *common.CivilTime `json:"end_date" gorm:"column:end_date"`
}

func (BookingUpdate) TableName() string {
	return Booking{}.TableName()
}

var (
	ErrInvalidRoomType = common.NewCustomError(
		nil,
		"room type is invalid",
		"ERR_INVALID_ROOM_TYPE")
)
