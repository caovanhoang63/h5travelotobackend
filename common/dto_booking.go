package common

import "time"

type DTOBooking struct {
	Id        int        `json:"id" gorm:"column:id"`
	HotelId   int        `json:"hotel_id" gorm:"column:hotel_id"`
	StartDate *time.Time `json:"start_date" gorm:"column:start_date"`
	EndDate   *time.Time `json:"end_date" gorm:"column:end_date"`
}
