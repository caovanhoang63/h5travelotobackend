package common

type DTORoom struct {
	Id         int `json:"id" gorm:"column:id;"`
	HotelId    int `json:"-" gorm:"column:hotel_id;"`
	RoomTypeID int `json:"-" gorm:"column:room_type_id;"`
	Status     int `json:"status" gorm:"column:status;"`
}

func (DTORoom) TableName() string {
	return "rooms"
}
