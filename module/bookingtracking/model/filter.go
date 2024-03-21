package bookingtrackingmodel

type Filter struct {
	HotelId int    `json:"hotel_id" form:"hotel_id"`
	State   string `json:"state" form:"state"`
}
