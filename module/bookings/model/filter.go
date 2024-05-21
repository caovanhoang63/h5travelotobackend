package bookingmodel

import (
	"h5travelotobackend/common"
)

type Filter struct {
	UserId     int               `json:"user_id" gorm:"column:user_id" form:"user_id" `
	HotelId    int               `json:"hotel_id" gorm:"column:hotel_id" form:"hotel_id"`
	RoomTypeID int               `json:"room_type_id" gorm:"column:room_type_id" form:"room_type_id"`
	StartDate  *common.CivilDate `json:"start_date" gorm:"column:start_date" form:"start_date"`
	EndDate    *common.CivilDate `json:"end_date" gorm:"end_date" form:"end_date"`
	State      string            `json:"state" gorm:"column:state" form:"state"`
	PayInHotel bool              `json:"pay_in_hotel" gorm:"column:pay_in_hotel" form:"pay_in_hotel"`
}
