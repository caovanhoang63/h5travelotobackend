package bookingbiz

import (
	"errors"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	bookingmodel "h5travelotobackend/module/bookings/model"
	bookingtrackingmodel "h5travelotobackend/module/bookingtracking/model"
	"log"
)

type UpdateStateStore interface {
	UpdateBookingState(ctx context.Context, id int, state string) error
	FindWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*bookingmodel.Booking, error)
}

type updateStateBiz struct {
	store UpdateStateStore
	pb    pubsub.Pubsub
}

func NewUpdateStateBiz(store UpdateStateStore, pb pubsub.Pubsub) *updateStateBiz {
	return &updateStateBiz{store: store, pb: pb}
}

func ValidateBookingStateUpdate(booking *bookingmodel.Booking, state string) error {
	if booking.State == common.BookingStateDeleted ||
		booking.State == common.BookingStateExpired ||
		booking.State == common.BookingStateCanceled {
		return common.ErrEntityDeleted(bookingmodel.EntityName, nil)
	}

	if booking.State == state {
		return bookingmodel.ErrStateAlreadyUpdate
	}

	return nil
}

func (u *updateStateBiz) UpdateState(ctx context.Context, bookingId int, state string) error {
	cond := map[string]interface{}{"id": bookingId}
	booking, err := u.store.FindWithCondition(ctx, cond)
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrEntityNotFound(bookingmodel.EntityName, err)
		}
		return common.ErrInternal(err)
	}

	if err = ValidateBookingStateUpdate(booking, state); err != nil {
		return err
	}

	err = u.store.UpdateBookingState(ctx, bookingId, state)

	if err != nil {
		return common.ErrCannotUpdateEntity(bookingmodel.EntityName, err)
	}

	tracking := bookingtrackingmodel.BookingTrackingCreate{BookingId: bookingId, State: state}
	tracking.Mask(false)
	mess := pubsub.NewMessage(&tracking)

	if err = u.pb.Publish(ctx, common.TopicUpdateBookingState, mess); err != nil {
		log.Println(err)
	}

	return nil
}
