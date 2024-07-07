package bookingmodel

type BookingStatistic struct {
	TotalCheckIn  int `json:"total_checked_in" gorm:"column:total_checked_in"`
	TotalCheckOut int `json:"total_checked_out" gorm:"column:total_checked_out"`
	TotalInHotel  int `json:"total_in_hotel" gorm:"column:total_in_hotel"`
}

type RoomStatus struct {
	Total     int `json:"total" gorm:"column:total"`
	Available int `json:"available" gorm:"column:available"`
	Booked    int `json:"booked" gorm:"column:booked"`
	Dirty     int `json:"dirty" gorm:"column:dirty"`
	Fixing    int `json:"fixing" gorm:"column:fixing"`
}
