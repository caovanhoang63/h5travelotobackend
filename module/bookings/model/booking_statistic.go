package bookingmodel

type BookingStatistic struct {
	CheckIn        int     `json:"check_in"`
	CheckOut       int     `json:"check_out"`
	InHotel        int     `json:"in_hotel"`
	AvailableRoom  int     `json:"available_room"`
	HotelTotalRoom int     `json:"hotel_total_room"`
	FloorStatus    float64 `json:"floor_status"`
}
