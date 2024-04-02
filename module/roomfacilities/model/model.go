package roomfacilitymodel

import "h5travelotobackend/common"

const EntityName = "RoomFacilities"

type RoomFacilityType struct {
	common.Facility `json:",inline"`
	Facilities      []RoomFacility `json:"facilities" gorm:"-"`
}

func (h *RoomFacilityType) Mask(isAdmin bool) {
	h.GenUID(common.DbTypeRoomFacilityType)
	if h.Facilities != nil {
		for i := range h.Facilities {
			h.Facilities[i].GenUID(common.DbTypeRoomFacility)
		}
	}
}

type RoomFacility common.FacilityDetail

func (RoomFacilityType) TableName() string {
	return "room_facility_types"
}

func (RoomFacility) TableName() string {
	return "room_facilities"
}

func (RoomFacilityType) CollectionName() string {
	return "room_facility_details"
}

type RoomFacilityDetail struct {
	RoomId     int `json:"Room_id" gorm:"column:Room_id"`
	FacilityId int `json:"facility_id" gorm:"column:facility_id"`
}

func (RoomFacilityDetail) TableName() string {
	return "Room_facility_details"
}
