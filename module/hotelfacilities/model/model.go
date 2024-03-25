package hotelfacilitymodel

import "h5travelotobackend/common"

const EntityName = "HotelFacilities"

type HotelFacility common.Facility

func (HotelFacility) CollectionName() string {
	return "hotel_facility_details"
}
