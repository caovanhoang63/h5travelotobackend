package payinbiz

import (
	"errors"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/payment/vnpay"
	payinmodel "h5travelotobackend/payment/module/payin/model"
)

type vnpIPNBiz struct {
	bkStore BookingStore
	pbStore UpdatPBStatusStore
	peStore PaymentEventStore
	vnPay   *vnpay.VnPay
}

func NewVnpIPNBiz(pbStore UpdatPBStatusStore, peStore PaymentEventStore, bkstore BookingStore, vnPay *vnpay.VnPay) *vnpIPNBiz {
	return &vnpIPNBiz{
		bkStore: bkstore,
		pbStore: pbStore,
		peStore: peStore,
		vnPay:   vnPay,
	}
}

func (biz *vnpIPNBiz) HandleIPNRequest(ctx context.Context, request *vnpay.IPNRequest) *vnpay.IPNResponse {
	bkId, err := request.GetBookingId()
	if err != nil {
		return vnpay.NewOrderNotFound()
	}
	bk, err := biz.bkStore.GetBookingById(ctx, bkId)
	if err != nil {
		return vnpay.NewOrderNotFound()
	}
	if bk.FinalAmount != request.GetAmount() {
		return vnpay.NewInvalidAmount()
	}

	pb, err := biz.pbStore.FindExecutingOrSuccessByBookingId(ctx, bkId)
	if err != nil && errors.Is(err, common.RecordNotFound) {
		return vnpay.NewOrderNotFound()
	}

	if pb.PaymentStatus == common.PaymentStatusSuccess || pb.PaymentStatus == common.PaymentStatusFailed {
		return vnpay.NewOrderAlreadyConfirmed()
	}

	if request.VnpResponseCode == vnpay.IPNSuccess {
		err = biz.pbStore.UpdateStatus(ctx, request.VnpTxnRef,
			&payinmodel.PaymentBookingUpdateStatus{PaymentStatus: common.PaymentStatusSuccess})
	} else {
		err = biz.pbStore.UpdateStatus(ctx, request.VnpTxnRef,
			&payinmodel.PaymentBookingUpdateStatus{PaymentStatus: common.PaymentStatusFailed})
	}

	if err != nil {
		return vnpay.NewOtherError()
	}

	return vnpay.NewSuccessResponse()
}
