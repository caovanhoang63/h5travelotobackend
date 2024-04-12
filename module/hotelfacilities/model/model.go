package hotelfacilitymodel

import "h5travelotobackend/common"

const EntityName = "HotelFacilities"

type HotelFacilityType struct {
	common.Facility `json:",inline"`
	Facilities      []HotelFacility `json:"facilities" gorm:"-"`
}

func (h *HotelFacilityType) Mask(isAdmin bool) {
	h.GenUID(common.DbTypeHotelFacilityType)
	if h.Facilities != nil {
		for i := range h.Facilities {
			h.Facilities[i].Mask(isAdmin)
		}
	}
}

type HotelFacility common.FacilityDetail

func (h *HotelFacility) Mask(isAdmin bool) {
	h.GenUID(common.DbTypeHotelFacility)
}

func (HotelFacilityType) TableName() string {
	return "hotel_facility_types"
}

func (HotelFacility) TableName() string {
	return "hotel_facilities"
}

func (HotelFacilityType) CollectionName() string {
	return "hotel_facility_details"
}

type HotelFacilityDetail struct {
	HotelId    int `json:"hotel_id" gorm:"column:hotel_id"`
	FacilityId int `json:"facility_id" gorm:"column:facility_id"`
}

func (HotelFacilityDetail) TableName() string {
	return "hotel_facility_details"
}
