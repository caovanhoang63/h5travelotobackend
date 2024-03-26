package common

import "time"

type DTOBooking struct {
	Id           int        `json:"id" gorm:"column:id"`
	HotelId      int        `json:"hotel_id" gorm:"column:hotel_id"`
	StartDate    *time.Time `json:"start_date" gorm:"column:start_date"`
	EndDate      *time.Time `json:"end_date" gorm:"column:end_date"`
	RoomQuantity int        `json:"room_quantity" gorm:"column:room_quantity"`
	RoomTypeId   int        `json:"-" gorm:"column:room_type_id"`
	Status       int        `json:"status" gorm:"column:status"`
}

func (DTOBooking) TableName() string {
	return "bookings"
}
