package bookingbiz

import (
	"context"
	"h5travelotobackend/common"
	"h5travelotobackend/module/bookings/model"
)

type FindBookingStore interface {
	FindWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*bookingmodel.Booking, error)
}

type findBookingBiz struct {
	store FindBookingStore
}

func NewFindBookingBiz(store FindBookingStore) *findBookingBiz {
	return &findBookingBiz{store: store}
}

func (biz *findBookingBiz) GetBookingById(ctx context.Context, id int) (*bookingmodel.Booking, error) {
	data, err := biz.store.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrEntityNotFound(bookingmodel.EntityName, err)
	}

	return data, nil
}
