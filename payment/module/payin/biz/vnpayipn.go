package payinbiz

import (
	"errors"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/payment/vnpay"
	payinmodel "h5travelotobackend/payment/module/payin/model"
)

type vnpIPNBiz struct {
	bkStore    BookingStore
	pbStore    UpdatPBStatusStore
	peStore    PaymentEventStore
	vnPay      *vnpay.VnPay
	successBiz *successPBStatusBiz
}

func NewVnpIPNBiz(pbStore UpdatPBStatusStore, peStore PaymentEventStore, bkstore BookingStore, vnPay *vnpay.VnPay, successBiz *successPBStatusBiz) *vnpIPNBiz {
	return &vnpIPNBiz{
		bkStore:    bkstore,
		pbStore:    pbStore,
		peStore:    peStore,
		vnPay:      vnPay,
		successBiz: successBiz,
	}
}

func (biz *vnpIPNBiz) HandleIPNRequest(ctx context.Context, request *vnpay.IPNRequest) *vnpay.IPNResponse {
	bkId, err := request.GetBookingId()
	if err != nil {
		return vnpay.NewOrderNotFound()
	}
	_, err = biz.bkStore.GetBookingById(ctx, bkId)
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return vnpay.NewOrderNotFound()
		}
		return vnpay.NewOtherError()
	}
	pb, err := biz.pbStore.FindWithCondition(ctx, map[string]interface{}{"txn_id": request.VnpTxnRef})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return vnpay.NewOrderNotFound()
		}
		return vnpay.NewOrderAlreadyConfirmed()
	}

	if pb.PaymentStatus == common.PaymentStatusSuccess || pb.PaymentStatus == common.PaymentStatusFailed {
		return vnpay.NewOrderAlreadyConfirmed()
	}

	if pb.Amount != request.GetAmount() {
		return vnpay.NewInvalidAmount()
	}

	if request.VnpResponseCode == vnpay.IPNSuccess {
		err = biz.successBiz.SuccessStatusBiz(ctx, pb)
	} else {
		err = biz.pbStore.UpdateStatus(ctx,
			pb.TxnId,
			&payinmodel.PaymentBookingUpdateStatus{
				PaymentStatus: common.PaymentStatusFailed})
	}
	if err != nil {
		return vnpay.NewOtherError()
	}
	return vnpay.NewSuccessResponse()
}
