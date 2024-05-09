package bookingmodel

import (
	"h5travelotobackend/common"
)

const EntityName = "Booking"

type Booking struct {
	common.SqlModel  `json:",inline"`
	HotelId          int                `json:"-" gorm:"column:hotel_id"`
	Hotel            *Hotel             `json:"hotel,omitempty" gorm:"foreignKey:HotelId;preload:false"`
	UserFakeId       *common.UID        `json:"user_id"`
	UserId           int                `json:"-" gorm:"column:user_id"`
	User             *common.SimpleUser `json:"user,omitempty" gorm:"foreignKey:UserId;preload:false"`
	RoomTypeId       int                `json:"-" gorm:"column:room_type_id"`
	RoomTypeFakeId   *common.UID        `json:"room_type_id" gorm:"-"`
	RoomQuantity     int                `json:"room_quantity" gorm:"column:room_quantity"`
	CustomerQuantity int                `json:"customer_quantity" gorm:"column:customer_quantity"`
	StartDate        *common.CivilDate  `json:"start_date" gorm:"column:start_date"`
	EndDate          *common.CivilDate  `json:"end_date" gorm:"column:end_date"`
}

func (Booking) TableName() string {
	return "bookings"
}

func (b *Booking) Mask(isAdmin bool) {
	b.GenUID(common.DbTypeBooking)
	b.UserFakeId = common.NewUIDP(uint32(b.UserId), common.DbTypeUser, 0)
	b.RoomTypeFakeId = common.NewUIDP(uint32(b.RoomTypeId), common.DbTypeRoomType, 0)
	if b.Hotel != nil {
		b.Hotel.Mask(isAdmin)
	}
	if b.User != nil {
		b.User.Mask(isAdmin)
	}
}

func (b *Booking) UnMask() {
	b.UserId = int(b.UserFakeId.GetLocalID())
	b.RoomTypeId = int(b.RoomTypeFakeId.GetLocalID())
}

type BookingCreate struct {
	common.SqlModel  `json:",inline"`
	HotelId          int               `json:"-" gorm:"column:hotel_id"`
	HotelFakeId      *common.UID       `json:"hotel_id" gorm:"-"`
	UserId           int               `json:"-" gorm:"column:user_id"`
	RoomTypeId       int               `json:"-" gorm:"column:room_type_id"`
	RoomTypeFakeId   *common.UID       `json:"room_type_id" gorm:"-"`
	CustomerQuantity int               `json:"customer_quantity" gorm:"column:customer_quantity"`
	RoomQuantity     int               `json:"room_quantity" gorm:"column:room_quantity"`
	StartDate        *common.CivilDate `json:"start_date" gorm:"column:start_date"`
	EndDate          *common.CivilDate `json:"end_date" gorm:"column:end_date"`
}

func (b *BookingCreate) UnMask() {
	b.HotelId = int(b.HotelFakeId.GetLocalID())
	b.RoomTypeId = int(b.RoomTypeFakeId.GetLocalID())
}

func (b *BookingCreate) Mask(isAdmin bool) {
	b.GenUID(common.DbTypeBooking)
	b.HotelFakeId = common.NewUIDP(uint32(b.HotelId), common.DbTypeHotel, 0)
	b.RoomTypeId = int(b.RoomTypeFakeId.GetLocalID())
}

func (BookingCreate) TableName() string {
	return Booking{}.TableName()
}

type BookingUpdate struct {
	CustomerQuantity int               `json:"customer_quantity" gorm:"column:customer_quantity"`
	StartDate        *common.CivilDate `json:"start_date" gorm:"column:start_date"`
	EndDate          *common.CivilDate `json:"end_date" gorm:"column:end_date"`
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

type Hotel struct {
	common.SqlModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name"`
	Address         string         `json:"address" gorm:"column:address"`
	HotelType       int            `json:"-" gorm:"column:hotel_type"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo"`
	Images          *common.Images `json:"images" gorm:"column:images"`
	ProvinceCode    int            `json:"-" gorm:"column:province_code"`
	Province        *Province      `json:"province,inline" gorm:"foreignKey:ProvinceCode;references:Code"`
	DistrictCode    int            `json:"-" gorm:"column:district_code"`
	District        *District      `json:"district,inline" gorm:"foreignKey:DistrictCode;references:Code"`
	WardCode        int            `json:"-" gorm:"column:ward_code"`
	Ward            *Ward          `json:"ward,inline" gorm:"foreignKey:WardCode;references:Code"`
	Star            int            `json:"star" gorm:"star"`
	TotalRating     int            `json:"total_rating" gorm:"total_rating"`
}

func (Hotel) TableName() string {
	return "hotels"
}

func (data *Hotel) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeHotel)
}

type Province struct {
	Code int    `json:"code" gorm:"column:code"`
	Name string `json:"name" gorm:"column:name"`
}

func (Province) TableName() string {
	return "provinces"
}

type District struct {
	Code int    `json:"code" gorm:"column:code"`
	Name string `json:"name" gorm:"column:name"`
}

func (District) TableName() string {
	return "districts"
}

type Ward struct {
	Code int    `json:"code" gorm:"column:code"`
	Name string `json:"name" gorm:"column:name"`
}

func (Ward) TableName() string {
	return "wards"
}
