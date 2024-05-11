package rtsearchmodel

import (
	"github.com/shopspring/decimal"
	"h5travelotobackend/common"
)

const EntityName = "RoomType"
const IndexName = "room_types_enriched"

type RoomType struct {
	common.SqlModel `json:",inline"`
	HotelId         int              `json:"hotel_id"`
	Name            string           `json:"name"`
	MaxCustomer     int              `json:"max_customer"`
	Area            int              `json:"area"`
	Bed             *common.Bed      `json:"bed"`
	Images          *common.Images   `json:"images"`
	BedStr          *string          `json:"bed_str"`
	ImagesStr       *string          `json:"images_str"`
	Price           *decimal.Decimal `json:"price"`
	TotalRoom       int              `json:"total_room"`
	PayInHotel      bool             `json:"pay_in_hotel"`
	BreakFast       bool             `json:"break_fast"`
	FreeCancel      bool             `json:"free_cancel"`
	AvailableRoom   int              `json:"available_room"`
	FacilityList    []int            `json:"facility_list"`
}
