package refundbiz

import (
	"errors"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/uuid"
	bookingmodel "h5travelotobackend/module/bookings/model"
	hoteldetailmodel "h5travelotobackend/module/hoteldetails/model"
	payinmodel "h5travelotobackend/payment/module/payin/model"
	paymenteventmodel "h5travelotobackend/payment/module/paymentevent/model"
	refundmodel "h5travelotobackend/payment/module/refund/model"
	"log"
)

type HotelDetailStore interface {
	GetHotelDetailById(ctx context.Context, id int) (*hoteldetailmodel.HotelDetail, error)
}

type BookingStore interface {
	GetBookingById(ctx context.Context, id int) (*bookingmodel.Booking, error)
}

type RefundStore interface {
	Create(ctx context.Context,
		create *refundmodel.RefundBookingCreate) error
}

type PEStore interface {
	Create(ctx context.Context, create *paymenteventmodel.PaymentEventCreate) error
}

type PBStore interface {
	FindWithCondition(ctx context.Context, cond map[string]interface{}) (*payinmodel.PaymentBooking, error)
}

type refundBiz struct {
	hStore  HotelDetailStore
	rStore  RefundStore
	peStore PEStore
	bkStore BookingStore
	pbStore PBStore
	uuid    uuid.Uuid
}

func NewRefundBiz(hStore HotelDetailStore, rStore RefundStore, peStore PEStore, bkStore BookingStore, pbStore PBStore, uuid uuid.Uuid) *refundBiz {
	return &refundBiz{hStore: hStore, rStore: rStore, peStore: peStore, bkStore: bkStore, pbStore: pbStore, uuid: uuid}
}

func (biz *refundBiz) Refund(ctx context.Context, refund *refundmodel.RefundBookingCreate) error {
	booking, err := biz.bkStore.GetBookingById(ctx, refund.BookingId)

	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.NewResourceNotFound(bookingmodel.EntityName)
		}
		return common.ErrInternal(err)
	}

	if booking.PayInHotel {
		return common.ErrInvalidRequest(errors.New("This booking is paid in hotel"))
	}

	if booking.Status == 0 {
		return common.ErrEntityDeleted(bookingmodel.EntityName, nil)
	}

	if booking.State != common.BookingStatePaid {
		return common.ErrInvalidRequest(errors.New("This booking is not paid"))
	}

	pb, err := biz.pbStore.FindWithCondition(ctx, map[string]interface{}{
		"booking_id":     refund.BookingId,
		"payment_status": common.PaymentStatusSuccess})

	log.Println(pb)
	if errors.Is(err, common.RecordNotFound) {
		return common.ErrInvalidRequest(errors.New("This booking is paid in hotel"))
	}

	hd, err := biz.hStore.GetHotelDetailById(ctx, booking.HotelId)
	if err != nil {
		return common.ErrInternal(err)
	}

	policy := hd.CancellationPolicy

	if policy == 1 {
		return nil
	}
	if policy == 0 {
		refund.Type = common.RefundTypeFull
	} else {
		refund.Type = common.RefundTypePartial
	}

	refund.PayInTxnId = pb.TxnId
	refund.Currency = pb.Currency
	refund.Amount = (1 - float64(policy)) * pb.Amount
	refund.CustomerId = pb.CustomerId
	refund.HotelId = pb.HotelId
	refund.PayInDate = pb.CreatedAt

	txn_id, _ := biz.uuid.Generate()
	pe := paymenteventmodel.PaymentEventCreate{
		TxnId:       txn_id,
		HotelId:     refund.HotelId,
		CustomerId:  refund.CustomerId,
		PaymentType: common.PaymentTypeRefund,
	}
	err = biz.peStore.Create(ctx, &pe)
	if err != nil {
		return paymenteventmodel.NewErrCannotCreatePaymentEvent(err)
	}

	refund.TxnId = pe.TxnId

	err = biz.rStore.Create(ctx, refund)
	if err != nil {
		return common.ErrInternal(err)
	}

	return nil
}
