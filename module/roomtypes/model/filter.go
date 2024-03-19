package roomtypemodel

import (
	"h5travelotobackend/common"
	"time"
)

type Filter struct {
	HotelId     int         `json:"-" gorm:"column:hotel_id;" form:"-"`
	HotelFakeId common.UID  `json:"hotel_id" gorm:"-" form:"hotel-id"`
	Bed         *common.Bed `json:"bed" gorm:"column:bed;" form:"bed"`

	BreakFast  bool    `json:"break_fast" gorm:"column:break_fast;" form:"break-fast"`
	FreeCancel bool    `json:"free_cancel" gorm:"column:free_cancel;" form:"free-cancel"`
	Rating     float64 `json:"rating" gorm:"column:rating;" form:"rating"`

	MaxPrice float64 `json:"max_price"  form:"max-price"`
	MinPrice float64 `json:"min_price" form:"min-price"`

	StartDate *time.Time `json:"start_date" form:"start-date"`
	EndDate   *time.Time `json:"end_date" form:"end-date"`
}

func (f *Filter) SetDefault() {
	f.MaxPrice = 10000000000
	f.MinPrice = 0
}

func (f *Filter) UnMask() {
	f.HotelId = int(f.HotelFakeId.GetLocalID())
}
