package common

type RoomTypeDTO struct {
	SqlModel         `json:",inline"`
	Name             string `json:"name" gorm:"column:name;"`
	HotelId          int    `json:"hotel_id" gorm:"column:hotel_id;"`
	CurAvailableRoom int    `json:"cur_available_room" gorm:"column:cur_available_room;"`
	TotalRoom        int    `json:"total_room" gorm:"column:total_room;"`
}

func (RoomTypeDTO) TableName() string {
	return "room_types"
}
