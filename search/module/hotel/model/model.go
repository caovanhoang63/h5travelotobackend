package hotelmodel

import (
	"h5travelotobackend/common"
)

const EntityName = "hotel"
const IndexName = "hotels_enriched"

type Hotel struct {
	common.ElasticModel
	OwnerID         int            `json:"-"`
	Name            string         `json:"name"`
	Address         string         `json:"address"`
	HotelType       int            `json:"-"`
	HotelTypeFakeId *common.UID    `json:"hotel_type"`
	Hotline         string         `json:"hotline"`
	Logo            *common.Image  `json:"logo"`
	Images          *common.Images `json:"images"`
	ProvinceCode    string         `json:"-"`
	Province        string         `json:"province"`
	District        string         `json:"district"`
	Ward            string         `json:"ward"`
	Lat             float64        `json:"lat"`
	Lng             float64        `json:"lng"`
	Star            int            `json:"star"`
	TotalRating     int            `json:"total_rating"`
	TotalRoomType   int            `json:"total_room_type"`
}
