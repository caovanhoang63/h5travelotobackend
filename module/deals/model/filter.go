package dealmodel

import (
	"h5travelotobackend/common"
	"time"
)

type Filter struct {
	HotelFakeId    *common.UID `json:"hotel_id" form:"hotel_id" gorm:"-"`
	HotelId        int         `json:"-" form:"-" gorm:"column:hotel_id;"`
	RoomTypeFakeId *common.UID `json:"room_type_id" form:"room_type_id" gorm:"-"`
	RoomTypeId     int         `json:"-" form:"-" gorm:"column:room_type_id;"`
	State          *string     `json:"state" form:"state" gorm:"column:state;"`
	StartDate      *time.Time  `json:"start_date" form:"start_date" gorm:"column:start_date;"`
	ExpiryDate     *time.Time  `json:"expiry_date" form:"expiry_date" gorm:"column:expiry_date;"`
}

func (f *Filter) Mask() {
	f.HotelFakeId = common.NewUIDP(uint32(f.HotelId), common.DbTypeHotel, 0)
	f.RoomTypeFakeId = common.NewUIDP(uint32(f.RoomTypeId), common.DbTypeRoomType, 0)
}

func (f *Filter) UnMask() {
	f.HotelId = int(f.HotelFakeId.GetLocalID())
	f.RoomTypeId = int(f.RoomTypeFakeId.GetLocalID())
}
