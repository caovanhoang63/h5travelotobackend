package bookingbiz

import (
	"context"
	"h5travelotobackend/common"
	"h5travelotobackend/module/bookings/bookingmodel"
)

type CreateBookingStore interface {
	Create(ctx context.Context, data *bookingmodel.BookingCreate) error
}

type createBookingBiz struct {
	store CreateBookingStore
}

func NewCreateBookingBiz(store CreateBookingStore) *createBookingBiz {
	return &createBookingBiz{store: store}
}

func (biz *createBookingBiz) Create(ctx context.Context, data *bookingmodel.BookingCreate) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(bookingmodel.EntityName, err)
	}
	return nil
}
