package roommodel

import "h5travelotobackend/common"

type Filter struct {
	HotelId        int         `json:"-" form:"-"`
	HotelFakeID    *common.UID `json:"hotel-id" form:"hotel-id"`
	RoomTypeId     int         `json:"-" form:"-"`
	RoomTypeFakeId *common.UID `json:"room-type-id" form:"room-type-id"`
	Status         int         `json:"status" form:"status"`
}

func (f *Filter) UnMask() *Filter {
	if f.HotelFakeID != nil {
		f.HotelId = int(f.HotelFakeID.GetLocalID())
	}
	if f.RoomTypeFakeId != nil {
		f.RoomTypeId = int(f.RoomTypeFakeId.GetLocalID())
	}
	return f

}

func (f *Filter) SetDefault() {
	f.Status = 1
}
