package payinbiz

import (
	"errors"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	payinmodel "h5travelotobackend/payment/module/payin/model"
)

type FindPaymentBookingStore interface {
	FindWithCondition(ctx context.Context, cond map[string]interface{}) (*payinmodel.PaymentBooking, error)
}

type findPaymentBookingBiz struct {
	pbStore FindPaymentBookingStore
}

func NewFindPaymentBookingBiz(pbStore FindPaymentBookingStore) *findPaymentBookingBiz {
	return &findPaymentBookingBiz{pbStore: pbStore}
}

func (biz *findPaymentBookingBiz) FindPBSuccessOfBooking(ctx context.Context, requester common.Requester, bookingId int) (*payinmodel.PaymentBooking, error) {
	cond := map[string]interface{}{
		"booking_id":     bookingId,
		"payment_status": common.PaymentStatusSuccess,
	}
	pb, err := biz.pbStore.FindWithCondition(ctx, cond)
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return nil, common.NewResourceNotFound(payinmodel.EntityName)
		}
		return nil, common.ErrInternal(err)
	}
	if requester.GetUserId() != pb.CustomerId && requester.GetRole() != common.RoleAdmin {
		return nil, common.ErrNoPermission(nil)
	}

	return pb, nil
}
