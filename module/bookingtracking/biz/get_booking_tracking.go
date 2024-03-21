package bookingtrackingbiz

import (
	"context"
	"h5travelotobackend/common"
	bookingtrackingmodel "h5travelotobackend/module/bookingtracking/model"
)

type FindBookingTrackingStore interface {
	GetBookingTracking(ctx context.Context, bookingId int) (*bookingtrackingmodel.BookingTracking, error)
}

type findBookingTrackingBiz struct {
	store FindBookingTrackingStore
}

func NewFindBookingTrackingBiz(store FindBookingTrackingStore) *findBookingTrackingBiz {
	return &findBookingTrackingBiz{store: store}
}

func (biz *findBookingTrackingBiz) GetBookingTracking(ctx context.Context, bookingId int) (*bookingtrackingmodel.BookingTracking, error) {
	data, err := biz.store.GetBookingTracking(ctx, bookingId)
	if err != nil {
		return nil, common.ErrEntityNotFound(bookingtrackingmodel.EntityName, err)
	}
	return data, nil
}
