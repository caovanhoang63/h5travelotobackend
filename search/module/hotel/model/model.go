package hotelmodel

const EntityName = "hotel"
const IndexName = "hotels_enriched"

type Hotel struct {
	OwnerID       int    `json:"-"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	HotelType     int    `json:"hotel_type"`
	Hotline       string `json:"hotline"`
	Star          int    `json:"star"`
	TotalRating   int    `json:"total_rating"`
	TotalRoomType int    `json:"total_room_type"`
}
