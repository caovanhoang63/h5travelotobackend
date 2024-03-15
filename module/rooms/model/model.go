package roommodel

import "h5travelotobackend/common"

const EntityName = "Room"

type Room struct {
	common.SqlModel `json:",inline"`
	HotelFakeId     *common.UID `json:"hotel_id" gorm:"-"`
	HotelId         int         `json:"-" gorm:"column:hotel_id;"`
	RoomTypeFakeId  *common.UID `json:"room_type_id" gorm:"-"`
	RoomTypeID      int         `json:"-" gorm:"column:room_type_id;"`
	Name            string      `json:"name" gorm:"column:name;"`
}

func (Room) TableName() string {
	return "rooms"
}

func (r *Room) Mask(isAdmin bool) {
	r.GenUID(common.DbTypeRoom)
	hotelUid := common.NewUID(uint32(r.HotelId), common.DbTypeHotel, 1)
	roomTypeUid := common.NewUID(uint32(r.RoomTypeID), common.DbTypeRoomType, 1)
	r.HotelFakeId = &hotelUid
	r.RoomTypeFakeId = &roomTypeUid
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
