package roommodel

type Filter struct {
	HotelId    int `json:"hotel_id" form:"hotel_id"`
	RoomTypeId int `json:"room_type_id" form:"room_type_id"`
}