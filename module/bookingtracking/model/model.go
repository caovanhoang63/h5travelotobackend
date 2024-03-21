package bookingtrackingmodel

import "h5travelotobackend/common"

const EntityName = "BookingTracking"

type BookingTracking struct {
	common.SqlModel `json:",inline"`
	BookingId       int        `json:"-" gorm:"column:booking_id"`
	BookingFakeId   common.UID `json:"booking_id" gorm:"-"`
	State           string     `json:"state" gorm:"column:state"`
}

func (BookingTracking) TableName() string {
	return "booking_trackings"
}

func (b *BookingTracking) Mask(isAdmin bool) {
	b.GenUID(common.DbTypeBookingTracking)
	b.BookingFakeId = common.NewUID(uint32(b.BookingId), common.DbTypeBookingTracking, 1)
}

func (b *BookingTracking) UnMask() {
	b.BookingId = int(b.BookingFakeId.GetLocalID())
}

type BookingTrackingCreate struct {
	common.SqlModel `json:",inline"`
	BookingId       int        `json:"-" gorm:"column:booking_id"`
	BookingFakeId   common.UID `json:"booking_id" gorm:"-"`
	State           string     `json:"state" gorm:"column:state"`
}

func (b *BookingTrackingCreate) UnMask() {
	b.BookingId = int(b.BookingFakeId.GetLocalID())
}

func (b *BookingTrackingCreate) Mask(isAdmin bool) {
	b.GenUID(common.DbTypeBookingTracking)
}

func (BookingTrackingCreate) TableName() string {
	return BookingTracking{}.TableName()
}

type BookingTrackingUpdate struct {
	State string `json:"state" gorm:"column:state"`
}

func (BookingTrackingUpdate) TableName() string {
	return BookingTracking{}.TableName()
}
