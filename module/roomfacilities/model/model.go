package roomfacilitymodel

import "h5travelotobackend/common"

type RoomFacility common.Facility

func (RoomFacility) CollectionName() string {
	return "room_facilities"
}
