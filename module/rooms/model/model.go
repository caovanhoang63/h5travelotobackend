package roommodel

import "h5travelotobackend/common"

const EntityName = "Room"

type Room struct {
	common.SqlModel `json:",inline"`
	HotelId         int    `json:"hotel_id" gorm:"column:hotel_id;"`
	RoomTypeID      int    `json:"room_type_id" gorm:"column:room_type_id;"`
	Name            string `json:"name" gorm:"column:name;"`
}

func (Room) TableName() string {
	return "rooms"
}

func (r *Room) Mask(isAdmin bool) {
	r.GenUID(common.DbTypeRoom)
}

type RoomCreate struct {
	common.SqlModel `json:",inline"`
	HotelId         int    `json:"hotel_id" gorm:"column:hotel_id;"`
	RoomTypeID      int    `json:"room_type_id" gorm:"column:room_type_id;"`
	Name            string `json:"name" gorm:"column:name;"`
}

func (RoomCreate) TableName() string {
	return Room{}.TableName()
}

func (r *RoomCreate) Mask(isAdmin bool) {
	r.GenUID(common.DbTypeRoom)
}

type RoomUpdate struct {
	RoomTypeID int    `json:"room_type_id" gorm:"column:room_type_id;"`
	Name       string `json:"name" gorm:"column:name;"`
}

func (RoomUpdate) TableName() string {
	return Room{}.TableName()
}
