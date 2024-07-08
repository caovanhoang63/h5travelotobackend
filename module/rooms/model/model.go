package roommodel

import (
	"h5travelotobackend/common"
)

const EntityName = "Room"

type RoomType struct {
	common.SqlModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
}

type Room struct {
	common.SqlModel `json:",inline"`
	HotelFakeId     *common.UID `json:"hotel_id" gorm:"-"`
	HotelId         int         `json:"-" gorm:"column:hotel_id;"`
	RoomTypeId      int         `json:"-" gorm:"column:room_type_id;"`
	RoomType        *RoomType   `json:"room_type" gorm:"foreignKey:RoomTypeId;preload:false"`
	Name            string      `json:"name" gorm:"column:name;"`
	Floor           int         `json:"floor" gorm:"column:floor;"`
	Condition       string      `json:"condition" gorm:"column:condition;"`
}

func (Room) TableName() string {
	return "rooms"
}

func (r *Room) Mask(isAdmin bool) {
	r.GenUID(common.DbTypeRoom)
	hotelUid := common.NewUID(uint32(r.HotelId), common.DbTypeHotel, 1)
	r.HotelFakeId = &hotelUid
	if r.RoomType != nil {
		r.RoomType.GenUID(common.DbTypeRoomType)
	}
}

type RoomCreate struct {
	common.SqlModel `json:",inline"`
	HotelId         int    `json:"-" gorm:"column:hotel_id;"`
	RoomTypeFakeId  string `json:"room_type_id" gorm:"-"`
	RoomTypeID      int    `json:"-" gorm:"column:room_type_id;"`
	Name            string `json:"name" gorm:"column:name;"`
}

func (RoomCreate) TableName() string {
	return Room{}.TableName()
}

func (r *RoomCreate) Mask(isAdmin bool) {
	r.GenUID(common.DbTypeRoom)
}

type RoomUpdate struct {
	RoomTypeFakeId string `json:"room_type_id" gorm:"-"`
	RoomTypeID     int    `json:"-" gorm:"column:room_type_id;"`
	Name           string `json:"name" gorm:"column:name;"`
	Status         int    `json:"status" gorm:"column:status"`
}

func (RoomUpdate) TableName() string {
	return Room{}.TableName()
}
