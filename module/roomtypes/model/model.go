package roomtypemodel

import "h5travelotobackend/common"

const EntityName = "RoomType"

type RoomType struct {
	common.SqlModel  `json:",inline"`
	HotelId          int            `json:"-" gorm:"column:hotel_id;"`
	HotelFakeId      common.UID     `json:"hotel_id" gorm:"-"`
	Name             string         `json:"name" gorm:"column:name;"`
	MaxCustomer      int            `json:"max_customer" gorm:"column:max_customer;"`
	Area             float64        `json:"area" gorm:"column:area;"`
	Bed              *common.Bed    `json:"bed" gorm:"column:bed;"`
	Price            float64        `json:"price" gorm:"column:price;"`
	CurAvailableRoom int            `json:"cur_available_room" gorm:"column:cur_available_room;"`
	Images           *common.Images `json:"images" gorm:"column:images;"`
	TotalRoom        int            `json:"total_room" gorm:"column:total_room;"`
	PayInHotel       bool           `json:"pay_in_hotel" gorm:"column:pay_in_hotel;"`
	BreakFast        bool           `json:"break_fast" gorm:"column:break_fast;"`
	FreeCancel       bool           `json:"free_cancel" gorm:"column:free_cancel;"`
	Rating           float64        `json:"rating" gorm:"column:rating;"`
}

func (RoomType) TableName() string {
	return "room_types"
}

func (r *RoomType) Mask(isAdmin bool) {
	r.GenUID(common.DbTypeRoomType)
	r.HotelFakeId = common.NewUID(uint32(r.HotelId), common.DbTypeHotel, 1)
}

type RoomTypeCreate struct {
	common.SqlModel `json:",inline"`
	HotelId         int         `json:"hotel_id" gorm:"column:hotel_id;"`
	Name            string      `json:"name" gorm:"column:name;"`
	MaxCustomer     int         `json:"max_customer" gorm:"column:max_customer;"`
	Area            float64     `json:"area" gorm:"column:area;"`
	Bed             *common.Bed `json:"bed" gorm:"column:bed;"`
	Price           float64     `json:"price" gorm:"column:price;"`
	PayInHotel      bool        `json:"pay_in_hotel" gorm:"column:pay_in_hotel;"`
	BreakFast       bool        `json:"break_fast" gorm:"column:break_fast;"`
	FreeCancel      bool        `json:"free_cancel" gorm:"column:free_cancel;"`
	FacilityIds     []string    `json:"facility_ids" gorm:"-"`
}

func (RoomTypeCreate) TableName() string {
	return RoomType{}.TableName()
}

func (r *RoomTypeCreate) Mask(isAdmin bool) {
	r.GenUID(common.DbTypeRoomType)
}

type RoomTypeUpdate struct {
	HotelId          int            `json:"hotel_id" gorm:"column:hotel_id;"`
	Name             string         `json:"name" gorm:"column:name;"`
	MaxCustomer      int            `json:"max_customer" gorm:"column:max_customer;"`
	Area             float64        `json:"area" gorm:"column:area;"`
	Bed              *common.Bed    `json:"bed" gorm:"column:bed;"`
	Price            float64        `json:"price" gorm:"column:price;"`
	CurAvailableRoom int            `json:"cur_available_room" gorm:"column:cur_available_room;"`
	Images           *common.Images `json:"images" gorm:"column:images;"`
	TotalRoom        int            `json:"total_room" gorm:"column:total_room;"`
	PayInHotel       bool           `json:"pay_in_hotel" gorm:"column:pay_in_hotel;"`
	BreakFast        bool           `json:"break_fast" gorm:"column:break_fast;"`
	FreeCancel       bool           `json:"free_cancel" gorm:"column:free_cancel;"`
	Rating           float64        `json:"rating" gorm:"column:rating;"`
}

func (RoomTypeUpdate) TableName() string {
	return RoomType{}.TableName()
}
