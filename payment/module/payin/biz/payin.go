package payinbiz

import (
	"golang.org/x/net/context"
	bookingmodel "h5travelotobackend/module/bookings/model"
)

type BookingStore interface {
	FindWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*bookingmodel.Booking, error)
}

type payInBiz struct {
	bkStore BookingStore
}

func NewPayInBiz(bkStore BookingStore) *payInBiz {
	return &payInBiz{bkStore: bkStore}
}
