package bookingtrackingbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	bookingtrackingmodel "h5travelotobackend/module/bookingtracking/model"
)

type GetBookingTrackingStore interface {
	GetBookingTrackings(ctx context.Context, bookingId int) ([]bookingtrackingmodel.BookingTracking, error)
}

type getBookingTrackingsBiz struct {
	store GetBookingTrackingStore
}

func NewGetBookingTrackingsBiz(store GetBookingTrackingStore) *getBookingTrackingsBiz {
	return &getBookingTrackingsBiz{store: store}
}

func (biz *getBookingTrackingsBiz) GetBookingTrackings(ctx context.Context, bookingId int) ([]bookingtrackingmodel.BookingTracking, error) {
	data, err := biz.store.GetBookingTrackings(ctx, bookingId)
	if err != nil {
		return nil, common.ErrEntityNotFound(bookingtrackingmodel.EntityName, err)
	}
	return data, nil
}
