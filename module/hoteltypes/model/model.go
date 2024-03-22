package hoteltypemodel

import "h5travelotobackend/common"

const EntityName = "HotelType"

type HotelType struct {
	common.SqlModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name"`
	Description     string `json:"description" gorm:"column:description"`
}

func (HotelType) TableName() string {
	return "hotel_types"
}

func (h *HotelType) Mask(isAdmin bool) {
	h.GenUID(common.DbTypeHotelType)
}

type HotelTypeCreate HotelType

func (HotelTypeCreate) TableName() string {
	return HotelType{}.TableName()
}

func (h *HotelTypeCreate) Mask(isAdmin bool) {
	h.GenUID(common.DbTypeHotelType)
}

type HotelTypeUpdate struct {
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

func (HotelTypeUpdate) TableName() string {
	return HotelType{}.TableName()
}
