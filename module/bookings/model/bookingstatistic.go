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

type OccupancyStatistic struct {
	Day0 float64 `json:"day_0" gorm:"column:day_0"`
	Day1 float64 `json:"day_1" gorm:"column:day_1"`
	Day2 float64 `json:"day_2" gorm:"column:day_2"`
	Day3 float64 `json:"day_3" gorm:"column:day_3"`
	Day4 float64 `json:"day_4" gorm:"column:day_4"`
	Day5 float64 `json:"day_5" gorm:"column:day_5"`
	Day6 float64 `json:"day_6" gorm:"column:day_6"`
}
