package hotelmodel

import "h5travelotobackend/common"

const EntityName = "hotel"
const IndexName = "hotels_enriched"

type Hotel struct {
	Id            int                        `json:"-"`
	FakeId        *common.UID                `json:"id"`
	OwnerID       int                        `json:"-"`
	Name          string                     `json:"name"`
	Address       string                     `json:"address"`
	HotelType     int                        `json:"hotel_type"`
	Hotline       string                     `json:"hotline"`
	Star          int                        `json:"star"`
	TotalRating   int                        `json:"total_rating"`
	TotalRoomType int                        `json:"total_room_type"`
	Location      *common.Location           `json:"location"`
	Province      *common.AdministrativeUnit `json:"province"`
	District      *common.AdministrativeUnit `json:"district"`
	Ward          *common.AdministrativeUnit `json:"ward"`
	Logo          *common.Image              `json:"logo"`
	Images        *common.Images             `json:"images"`
}

func (h *Hotel) Mask(isAdmin bool) {
	h.FakeId = common.NewUIDP(uint32(h.Id), common.DbTypeHotel, 0)
}

type HotelImage struct {
	LogoStr   *string `json:"logo_str"`
	ImagesStr *string `json:"images_str"`
}
