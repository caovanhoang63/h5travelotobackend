package payinbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	payinmodel "h5travelotobackend/payment/module/payin/model"
)

type UpdatPBStatusStore interface {
	FindExecutingOrSuccessByBookingId(ctx context.Context, bookingId int) (*payinmodel.PaymentBooking, error)
	UpdateStatus(ctx context.Context, txnId string, status *payinmodel.PaymentBookingUpdateStatus) error
}

type updatePBStatusBiz struct {
	pbStore UpdatPBStatusStore
}

func NewUpdatePBStatusBiz(pbStore UpdatPBStatusStore) *updatePBStatusBiz {
	return &updatePBStatusBiz{pbStore: pbStore}
}

func (biz *updatePBStatusBiz) UpdateStatus(ctx context.Context, txnId string, status *payinmodel.PaymentBookingUpdateStatus) error {
	err := biz.pbStore.UpdateStatus(ctx, txnId, status)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
