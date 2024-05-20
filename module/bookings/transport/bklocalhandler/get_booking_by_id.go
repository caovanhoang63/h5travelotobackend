package bklocalhandler

import (
	"golang.org/x/net/context"
	bookingbiz "h5travelotobackend/module/bookings/biz"
	bookingmodel "h5travelotobackend/module/bookings/model"
	bookingsqlstorage "h5travelotobackend/module/bookings/storage"
)

func (h *CountBookedRoomHandler) GetBookingById(ctx context.Context, id int) (*bookingmodel.Booking, error) {
	store := bookingsqlstorage.NewSqlStore(h.appCtx.GetGormDbConnection())
	biz := bookingbiz.NewFindBookingBiz(store)
	data, err := biz.GetBookingById(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
