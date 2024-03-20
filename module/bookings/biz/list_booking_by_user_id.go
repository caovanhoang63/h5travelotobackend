package bookingbiz

import (
	"context"
	"errors"
	"h5travelotobackend/common"
	"h5travelotobackend/module/bookings/bookingmodel"
)

type ListBookingStore interface {
	ListBookingWithCondition(
		ctx context.Context,
		filter *bookingmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]bookingmodel.Booking, error)
}

type listBookingBiz struct {
	store ListBookingStore
}

func NewListBookingBiz(store ListBookingStore) *listBookingBiz {
	return &listBookingBiz{store: store}
}

func (biz *listBookingBiz) ListBookingByUserId(
	ctx context.Context,
	requester common.Requester,
	userId int,
	filter *bookingmodel.Filter,
	paging *common.Paging,
) ([]bookingmodel.Booking, error) {
	if requester.GetUserId() != userId {
		return nil, common.ErrNoPermission(errors.New("user does not have permission to view this booking"))
	}
	filter.UserId = userId

	result, err := biz.store.ListBookingWithCondition(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(bookingmodel.EntityName, err)
	}

	return result, nil
}
