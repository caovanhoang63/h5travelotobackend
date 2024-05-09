package hotelmodel

import "h5travelotobackend/common"

const EntityName = "hotel"
const IndexName = "hotels_enriched"

type Hotel struct {
	Id            int              `json:"-"`
	FakeId        *common.UID      `json:"id"`
	OwnerID       int              `json:"-"`
	Name          string           `json:"name"`
	Address       string           `json:"address"`
	Logo          *common.Image    `json:"logo"`
	Images        *common.Images   `json:"images"`
	HotelType     int              `json:"hotel_type"`
	Hotline       string           `json:"hotline"`
	Star          int              `json:"star"`
	TotalRating   int              `json:"total_rating"`
	TotalRoomType int              `json:"total_room_type"`
	Location      *common.Location `json:"location"`
	Province      *Province        `json:"province"`
	District      *District        `json:"district"`
	Ward          *Ward            `json:"ward"`
}

type Province struct {
	Code string `json:"province_code"`
	Name string `json:"province_name"`
}

type District struct {
	Code string `json:"district_code"`
	Name string `json:"district_name"`
}

type Ward struct {
	Code string `json:"ward_code"`
	Name string `json:"ward_name"`
}

func (h *Hotel) Mask(isAdmin bool) {
	h.FakeId = common.NewUIDP(uint32(h.Id), common.DbTypeHotel, 0)
}

type HotelImage struct {
	LogoStr   *string `json:"logo_str"`
	ImagesStr *string `json:"images_str"`
}
