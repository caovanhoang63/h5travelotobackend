package htsavemodel

import (
	"h5travelotobackend/common"
	"time"
)

const EntityName = "HotelSave"

type HotelSave struct {
	UserId      int           `json:"user_id" gorm:"column:user_id"`
	UserFakeId  *common.UID   `json:"-" form:"-" gorm:"column:user_id;"`
	HotelId     int           `json:"-" gorm:"column:hotel_id"`
	Hotel       *common.Hotel `json:"hotel,omitempty" gorm:"foreignKey:HotelId;preload:false"`
	HotelFakeId *common.UID   `json:"-" form:"-" gorm:"-"`
	CreatedAt   *time.Time    `json:"created_at" gorm:"column:created_at"`
}

func (HotelSave) TableName() string {
	return "hotels_saved"
}

func (h *HotelSave) Mask() {
	h.UserFakeId = common.NewUIDP(uint32(h.UserId), common.DbTypeUser, 0)
	h.HotelFakeId = common.NewUIDP(uint32(h.HotelId), common.DbTypeHotel, 0)
}

type HotelSaveCreate struct {
	UserId  int `json:"user_id" form:"user_id" gorm:"column:user_id"`
	HotelId int `json:"hotel_id" form:"hotel_id" gorm:"column:hotel_id"`
}

func (h *HotelSaveCreate) TableName() string {
	return HotelSave{}.TableName()
}

type HotelSaveDelete struct {
	UserId  int `json:"user_id" form:"user_id" gorm:"column:user_id"`
	HotelId int `json:"hotel_id" form:"hotel_id" gorm:"column:hotel_id"`
}

func (h *HotelSaveDelete) TableName() string {
	return HotelSave{}.TableName()
}
func ErrCannotSaveHotel(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot save hotel",
		"Err_Cannot_Save_Hotel",
	)
}

func ErrCannotUnsaveHotel(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot unsave hotel",
		"Err_Cannot_Unsave_Hotel",
	)
}
