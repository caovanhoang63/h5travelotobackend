package payinbiz

import (
	"errors"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	payinmodel "h5travelotobackend/payment/module/payin/model"
)

type UpdatPBStatusStore interface {
	FindExecutingOrSuccessByBookingId(ctx context.Context, bookingId int) (*payinmodel.PaymentBooking, error)
	UpdateStatus(ctx context.Context, txnId string, status *payinmodel.PaymentBookingUpdateStatus) error
	FindWithCondition(ctx context.Context, cond map[string]interface{}) (*payinmodel.PaymentBooking, error)
}

type updatePBStatusBiz struct {
	pbStore UpdatPBStatusStore
}

func NewUpdatePBStatusBiz(pbStore UpdatPBStatusStore) *updatePBStatusBiz {
	return &updatePBStatusBiz{pbStore: pbStore}
}

func (biz *updatePBStatusBiz) UpdateStatus(ctx context.Context, requester common.Requester, txnId string, status *payinmodel.PaymentBookingUpdateStatus) error {
	pb, err := biz.pbStore.FindWithCondition(ctx, map[string]interface{}{"txn_id": txnId})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return payinmodel.ErrCannotFindTransaction
		}
		return common.ErrInternal(err)
	}

	if requester.GetUserId() != pb.CustomerId && requester.GetRole() != common.RoleAdmin {
		return common.ErrNoPermission(nil)
	}

	if pb.PaymentStatus == status.PaymentStatus {
		return payinmodel.ErrPaymentStatusAlreadyUpdated
	}

	err = biz.pbStore.UpdateStatus(ctx, txnId, status)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
