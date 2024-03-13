package roomtypemodel

import "h5travelotobackend/common"

type Filter struct {
	HotelId int         `json:"hotel_id" gorm:"column:hotel_id;" form:"hotel_id"`
	Bed     *common.Bed `json:"bed" gorm:"column:bed;" form:"bed"`

	BreakFast  bool    `json:"break_fast" gorm:"column:break_fast;" form:"break_fast"`
	FreeCancel bool    `json:"free_cancel" gorm:"column:free_cancel;" form:"free_cancel"`
	Rating     float64 `json:"rating" gorm:"column:rating;" form:"rating"`

	MaxPrice float64 `json:"price"  form:"max_price"`
	MinPrice float64 `json:"min_price" form:"min_price"`
}
