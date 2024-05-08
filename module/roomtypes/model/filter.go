package roomtypemodel

import (
	"h5travelotobackend/common"
)

type Filter struct {
	HotelId     int         `json:"-" gorm:"column:hotel_id;" form:"-"`
	HotelFakeId common.UID  `json:"hotel_id" gorm:"-" form:"hotel_id"`
	Bed         *common.Bed `json:"bed" gorm:"column:bed;" form:"bed"`
	Area        float64     `json:"area" gorm:"column:area;" form:"area"`
	PayInHotel  bool        `json:"pay_in_hotel" gorm:"column:pay_in_hotel;" form:"pay_in_hotel"`
	BreakFast   bool        `json:"break_fast" gorm:"column:break_fast;" form:"break_fast"`
	FreeCancel  bool        `json:"free_cancel" gorm:"column:free_cancel;" form:"free_cancel"`
	MaxPrice    float64     `json:"max_price"  form:"max_price" `
	MinPrice    float64     `json:"min_price" form:"min_price" `
	MaxCustomer int         `json:"max_customer" gorm:"column:max_customer;" form:"max_customer"`
}

func (f *Filter) SetDefault() {
	if f.MaxPrice == 0 {
		f.MaxPrice = 10000000000
	}
}

func (f *Filter) UnMask() {
	f.HotelId = int(f.HotelFakeId.GetLocalID())
}
