package refundbiz

import (
	"golang.org/x/net/context"
	hoteldetailmodel "h5travelotobackend/module/hoteldetails/model"
	paymenteventmodel "h5travelotobackend/payment/module/paymentevent/model"
	refundmodel "h5travelotobackend/payment/module/refund/model"
)

type HotelDetailStore interface {
	GetHotelDetailById(ctx context.Context, id int) (*hoteldetailmodel.HotelDetail, error)
}

type RefundStore interface {
	Create(ctx context.Context,
		create *refundmodel.RefundBookingCreate) error
}

type PEStore interface {
	Create(ctx context.Context, create *paymenteventmodel.PaymentEventCreate) error
}

type refundBiz struct {
	hStore  HotelDetailStore
	rStore  RefundStore
	peStore PEStore
}

func NewRefundBiz(hStore HotelDetailStore, rStore RefundStore, peStore PEStore) *refundBiz {
	return &refundBiz{hStore: hStore, rStore: rStore, peStore: peStore}
}

func (biz *refundBiz) Refund(ctx context.Context, refund *refundmodel.RefundBookingCreate) error {

	return nil
}
