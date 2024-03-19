package bookingbiz

import (
	"context"
	"errors"
	"h5travelotobackend/common"
	"h5travelotobackend/module/bookings/bookingmodel"
	workermodel "h5travelotobackend/module/worker/model"
)

type DeleteBookingStore interface {
	Delete(ctx context.Context, id int) error
}

type FindWorkerStore interface {
	FindWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*workermodel.Worker, error)
}

type deleteBookingBiz struct {
	deleteStore     DeleteBookingStore
	findStore       FindBookingStore
	findWorkerStore FindWorkerStore
}

func NewDeleteBookingBiz(deleteStore DeleteBookingStore,
	findStore FindBookingStore,
	findWorkerStore FindWorkerStore) *deleteBookingBiz {
	return &deleteBookingBiz{deleteStore: deleteStore, findStore: findStore, findWorkerStore: findWorkerStore}
}

func (biz *deleteBookingBiz) DeleteBooking(ctx context.Context, requester common.Requester, id int) error {

	oldData, err := biz.findStore.FindWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotDeleteEntity(bookingmodel.EntityName, err)
	}
	if oldData.Status == 0 {
		return common.ErrEntityDeleted(bookingmodel.EntityName, nil)
	}

	if requester.GetRole() == common.RoleCustomer {
		if requester.GetUserId() != oldData.UserId {
			return common.ErrNoPermission(errors.New("user does not have permission to delete this booking"))
		}
	} else if requester.GetRole() != common.RoleAdmin {
		worker, err := biz.findWorkerStore.FindWithCondition(ctx, map[string]interface{}{"user_id": requester.GetUserId()})
		if err != nil {
			return common.ErrNoPermission(err)
		}
		if worker.HotelId != oldData.HotelId {
			return common.ErrNoPermission(errors.New("worker is not in the same hotel as the booking"))
		}
	}

	if err := biz.deleteStore.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(bookingmodel.EntityName, err)
	}
	return nil
}
