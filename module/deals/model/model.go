package dealmodel

import (
	"h5travelotobackend/common"
	"time"
)

const EntityName = "Deal"

type Deal struct {
	common.SqlModel   `json:",inline"`
	Name              string        `json:"name" gorm:"column:name;"`
	HotelId           int           `json:"-" gorm:"column:hotel_id;"`
	HotelFakeId       *common.UID   `json:"hotel_id" gorm:"-"`
	RoomTypeId        int           `json:"-" gorm:"column:room_type_id;"`
	RoomTypeFakeId    *common.UID   `json:"room_type_id" gorm:"-"`
	Image             *common.Image `json:"image" gorm:"column:image;"`
	Description       string        `json:"description" gorm:"column:description;"`
	TotalQuantity     int           `json:"total_quantity" gorm:"column:total_quantity;"`
	AvailableQuantity int           `json:"available_quantity" gorm:"column:available_quantity;"`
	DiscountType      string        `json:"discount_type" gorm:"column:discount_type;"`
	DiscountAmount    float64       `json:"discount_amount" gorm:"column:discount_amount;"`
	DiscountPercent   float64       `json:"discount_percent" gorm:"column:discount_percent;"`
	StartDate         *time.Time    `json:"start_date" gorm:"column:start_date;"`
	ExpiryDate        *time.Time    `json:"expiry_date" gorm:"column:expiry_date;"`
	State             string        `json:"state" gorm:"column:state;"` // active, inactive
}

func (Deal) TableName() string {
	return "deals"
}

func (d *Deal) Mask(isAdmin bool) {
	d.GenUID(common.DbTypeDeal)
	d.HotelFakeId = common.NewUIDP(uint32(d.HotelId), common.DbTypeHotel, 0)
	d.RoomTypeFakeId = common.NewUIDP(uint32(d.RoomTypeId), common.DbTypeRoomType, 0)
}

func (d *Deal) UnMask() {
	d.HotelId = int(d.HotelFakeId.GetLocalID())
	d.RoomTypeId = int(d.RoomTypeFakeId.GetLocalID())
}

type DealCreate Deal // for create

func (DealCreate) TableName() string {
	return Deal{}.TableName()
}

func (d *DealCreate) Mask(isAdmin bool) {
	d.GenUID(common.DbTypeDeal)
	d.HotelFakeId = common.NewUIDP(uint32(d.HotelId), common.DbTypeHotel, 0)
	d.RoomTypeFakeId = common.NewUIDP(uint32(d.RoomTypeId), common.DbTypeRoomType, 0)
}
func (d *DealCreate) UnMask() {
	d.HotelId = int(d.HotelFakeId.GetLocalID())
	d.RoomTypeId = int(d.RoomTypeFakeId.GetLocalID())
}

type DealUpdate struct {
	Name              string        `json:"name" gorm:"column:name;"`
	HotelId           int           `json:"-" gorm:"column:hotel_id;"`
	HotelFakeId       *common.UID   `json:"hotel_id" gorm:"-"`
	RoomTypeId        int           `json:"-" gorm:"column:room_type_id;"`
	RoomTypeFakeId    *common.UID   `json:"room_type_id" gorm:"-"`
	Image             *common.Image `json:"image" gorm:"column:image;"`
	Description       string        `json:"description" gorm:"column:description;"`
	TotalQuantity     int           `json:"total_quantity" gorm:"column:total_quantity;"`
	AvailableQuantity int           `json:"available_quantity" gorm:"column:available_quantity;"`
	DiscountType      string        `json:"discount_type" gorm:"column:discount_type;"`
	DiscountAmount    float64       `json:"discount_amount" gorm:"column:discount_amount;"`
	DiscountPercent   float64       `json:"discount_percent" gorm:"column:discount_percent;"`
	StartDate         *time.Time    `json:"start_date" gorm:"column:start_date;"`
	ExpiryDate        *time.Time    `json:"expiry_date" gorm:"column:expiry_date;"`
	State             string        `json:"state" gorm:"column:state;"` // active, inactive
}

func (DealUpdate) TableName() string {
	return Deal{}.TableName()
}

func (d *DealUpdate) Mask(isAdmin bool) {
	d.HotelFakeId = common.NewUIDP(uint32(d.HotelId), common.DbTypeHotel, 0)
	d.RoomTypeFakeId = common.NewUIDP(uint32(d.RoomTypeId), common.DbTypeRoomType, 0)
}

func (d *DealUpdate) UnMask() {
	d.HotelId = int(d.HotelFakeId.GetLocalID())
	d.RoomTypeId = int(d.RoomTypeFakeId.GetLocalID())
}
