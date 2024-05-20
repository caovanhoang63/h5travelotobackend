package payinbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	bookingmodel "h5travelotobackend/module/bookings/model"
	payinmodel "h5travelotobackend/payment/module/payin/model"
	paymenteventmodel "h5travelotobackend/payment/module/paymentevent/model"
)

type BookingStore interface {
	GetBookingById(ctx context.Context, id int) (*bookingmodel.Booking, error)
	UpdateDeal(ctx context.Context, id, dealId int) error
}

type PaymentBookingStore interface {
	Create(ctx context.Context, create *payinmodel.PaymentBookingCreate) error
}

type PaymentEventStore interface {
	Create(ctx context.Context, create *paymenteventmodel.PaymentEventCreate) error
}

type payInBiz struct {
	bkStore BookingStore
	pbStore PaymentBookingStore
	peStore PaymentEventStore
}

func NewPayInBiz(pbStore PaymentBookingStore, peStore PaymentEventStore, bkstore BookingStore) *payInBiz {
	return &payInBiz{
		bkStore: bkstore,
		pbStore: pbStore,
		peStore: peStore,
	}
}

func (biz *payInBiz) NewPaymentBooking(ctx context.Context, requester common.Requester, create *payinmodel.PaymentBookingCreate) error {
	if create.DealId != 0 {
		err := biz.bkStore.UpdateDeal(ctx, create.BookingId, create.DealId)
		if err != nil {
			return common.ErrInvalidRequest(err)
		}
	}

	booking, err := biz.bkStore.GetBookingById(ctx, create.BookingId)
	if err != nil {
		return common.ErrInvalidRequest(err)
	}

	peCreate := paymenteventmodel.PaymentEventCreate{
		CustomerId:  requester.GetUserId(),
		HotelId:     booking.HotelId,
		PaymentType: common.PaymentTypePayIn,
	}

	if err = biz.peStore.Create(ctx, &peCreate); err != nil {
		return common.ErrCannotCreateEntity("Transaction", err)
	}

	create.Amount = booking.FinalAmount
	create.TxnId = peCreate.TxnId

	if err = biz.pbStore.Create(ctx, create); err != nil {
		return common.ErrCannotCreateEntity("Transaction", err)
	}

	return nil
}
