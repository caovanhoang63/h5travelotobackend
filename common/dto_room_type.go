package common

type DTORoomType struct {
	Id               int    `json:"id" gorm:"column:id;"`
	Name             string `json:"name" gorm:"column:name;"`
	HotelId          int    `json:"hotel_id" gorm:"column:hotel_id;"`
	CurAvailableRoom int    `json:"cur_available_room" gorm:"column:cur_available_room;"`
	TotalRoom        int    `json:"total_room" gorm:"column:total_room;"`
	Status           int    `json:"status" gorm:"column:status;"`
}

func (DTORoomType) TableName() string {
	return "room_types"
}
