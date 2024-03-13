package roomtypemodel

import (
	"h5travelotobackend/common"
	"time"
)

type Filter struct {
	HotelId int         `json:"hotel_id" gorm:"column:hotel_id;" form:"hotel_id"`
	Bed     *common.Bed `json:"bed" gorm:"column:bed;" form:"bed"`

	BreakFast  bool    `json:"break_fast" gorm:"column:break_fast;" form:"break_fast"`
	FreeCancel bool    `json:"free_cancel" gorm:"column:free_cancel;" form:"free_cancel"`
	Rating     float64 `json:"rating" gorm:"column:rating;" form:"rating"`

	MaxPrice float64 `json:"max_price"  form:"max_price"`
	MinPrice float64 `json:"min_price" form:"min_price"`

	StartDate *time.Time `json:"start_date" form:"start_date"`
	EndDate   *time.Time `json:"end_date" form:"end_date"`
}

func (f *Filter) SetDefault() {
	f.MaxPrice = 10000000000
	f.MinPrice = 0
}
