package bookingmodel

type BookingStatistic struct {
	TotalCheckIn  int `json:"total_checked_in" gorm:"column:total_checked_in"`
	TotalCheckOut int `json:"total_checked_out" gorm:"column:total_checked_out"`
	TotalInHotel  int `json:"total_in_hotel" gorm:"column:total_in_hotel"`
}
