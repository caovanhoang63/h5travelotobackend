package rtsearchmodel

import (
	"h5travelotobackend/common"
)

const EntityName = "RoomType"
const IndexName = "room_types_enriched"

type RoomType struct {
	common.SqlModel `json:",inline"`
	Name            string         `json:"name"`
	MaxCustomer     int            `json:"max_customer"`
	Area            *float64       `json:"area"`
	Bed             *common.Bed    `json:"bedJson"`
	Images          *common.Images `json:"imagesJson"`
	Price           *float64       `json:"price"`
	TotalRoom       int            `json:"total_room"`
	PayInHotel      int            `json:"pay_in_hotel"`
	BreakFast       int            `json:"break_fast"`
	FreeCancel      int            `json:"free_cancel"`
	AvailableRoom   int            `json:"available_room"`
	FacilityList    []int          `json:"facility_list"`
}

func (rt *RoomType) Mask(isAdmin bool) {
	rt.FakeId = common.NewUIDP(uint32(rt.Id), common.DbTypeRoomType, 0)
}

type RoomTypeStrFields struct {
	Bed    *string `json:"bed"`
	Images *string `json:"images"`
}
