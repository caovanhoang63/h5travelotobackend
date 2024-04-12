package invoicemodel

import "h5travelotobackend/common"

type Filter struct {
	BookingId     int               `json:"-" form:"-" gorm:"column:booking_id"`
	BookingFakeId *common.UID       `json:"booking_id" form:"booking_id" gorm:"column:booking_id"`
	DealId        int               `json:"-" form:"-" gorm:"column:deal_id"`
	DealFakeId    *common.UID       `json:"deal_id" form:"deal_id" gorm:"column:deal_id"`
	CreatedAt     *common.CivilTime `json:"created_at" form:"created_at" gorm:"column:created_at"`
}

func (f *Filter) Mask(isAdmin bool) {
	if f.BookingId != 0 {
		f.BookingFakeId = common.NewUIDP(uint32(f.BookingId), common.DbTypeBooking, 0)
	}
	if f.DealId != 0 {
		f.DealFakeId = common.NewUIDP(uint32(f.DealId), common.DbTypeDeal, 0)
	}
}

func (f *Filter) UnMask() {
	if f.BookingFakeId != nil {
		f.BookingId = int(f.BookingFakeId.GetLocalID())
	}
	if f.DealFakeId != nil {
		f.DealId = int(f.DealFakeId.GetLocalID())
	}
}
