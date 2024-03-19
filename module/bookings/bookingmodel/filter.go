package bookingmodel

import "time"

type Filter struct {
	UserId    int        `json:"user_id" gorm:"column:user_id"`
	HotelId   int        `json:"hotel_id" gorm:"column:hotel_id"`
	StartDate *time.Time `json:"start_date" gorm:"column:start_date"`
	EndDate   *time.Time `json:"end_date" gorm:"end_date"`
}
