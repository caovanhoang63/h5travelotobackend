package hotelmodel

import "h5travelotobackend/common"

const EntityName = "hotel"
const IndexName = "hotels_enriched"

type Hotel struct {
	Id            int            `json:"-"`
	FakeId        *common.UID    `json:"id"`
	OwnerID       int            `json:"-"`
	Name          string         `json:"name"`
	Address       string         `json:"address"`
	LogoStr       *string        `json:"logo_str"`
	ImagesStr     *string        `json:"images_str"`
	Logo          *common.Image  `json:"logo"`
	Images        *common.Images `json:"images"`
	HotelType     int            `json:"hotel_type"`
	Hotline       string         `json:"hotline"`
	Star          int            `json:"star"`
	TotalRating   int            `json:"total_rating"`
	TotalRoomType int            `json:"total_room_type"`
}
