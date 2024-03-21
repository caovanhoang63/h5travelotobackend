package bookingtrackingbiz

import (
	"context"
	"h5travelotobackend/common"
	bookingtrackingmodel "h5travelotobackend/module/bookingtracking/model"
)

type UpdateTrackingStore interface {
	Update(ctx context.Context, bookingId int, data *bookingtrackingmodel.BookingTrackingUpdate) error
}

type updateTrackingBiz struct {
	store UpdateTrackingStore
}

func NewUpdateTrackingBiz(store UpdateTrackingStore) *updateTrackingBiz {
	return &updateTrackingBiz{store: store}
}

func (biz *updateTrackingBiz) UpdateTracking(ctx context.Context, bookingId int, data *bookingtrackingmodel.BookingTrackingUpdate) error {
	if err := biz.store.Update(ctx, bookingId, data); err != nil {
		return common.ErrCannotUpdateEntity(bookingtrackingmodel.EntityName, err)
	}
	return nil
}
