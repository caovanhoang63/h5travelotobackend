package localrefund

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/payment/vnpay"
	"h5travelotobackend/module/bookings/transport/bklocalhandler"
	"h5travelotobackend/module/hoteldetails/transport/localhoteldetail"
	payinstore "h5travelotobackend/payment/module/payin/store"
	paymenteventstore "h5travelotobackend/payment/module/paymentevent/store"
	refundbiz "h5travelotobackend/payment/module/refund/biz"
	refundmodel "h5travelotobackend/payment/module/refund/model"
	refundstore "h5travelotobackend/payment/module/refund/store"
	"log"
)

type VnPayRefund struct {
	appCtx appContext.AppContext
}

func NewVnPayRefund(appCtx appContext.AppContext) *VnPayRefund {
	return &VnPayRefund{appCtx: appCtx}
}

func (v *VnPayRefund) VnPayRefund(ctx context.Context, reason string, bookingId int) error {
	create := &refundmodel.RefundBookingCreate{
		BookingId: bookingId,
		Method:    common.PaymentMethodVnPay,
		Reason:    reason,
	}

	vnp := v.appCtx.GetVnPay()

	hStore := localhoteldetail.NewHotelDetailLocalHandler(v.appCtx)
	rStore := refundstore.NewStore(v.appCtx.GetGormDbConnection())
	peStore := paymenteventstore.NewStore(v.appCtx.GetGormDbConnection())
	bkStore := bklocalhandler.NewCountBookedRoomLocalHandler(v.appCtx)
	pbStore := payinstore.NewStore(v.appCtx.GetGormDbConnection())
	uuid := v.appCtx.GetUUID()
	biz := refundbiz.NewRefundBiz(hStore, rStore, peStore, bkStore, pbStore, uuid)

	err := biz.Refund(ctx, create)
	if err != nil {
		return err
	}

	create.Mask(false)

	var refundType string
	if create.Type == common.RefundTypeFull {
		refundType = vnpay.RefundTypeFull
	} else {
		refundType = vnpay.RefundTypePart
	}

	url := vnp.NewRefundUrl(create.TxnId,
		create.PayInTxnId,
		create.BookingFakeId.String(),
		create.UserFakeId.String(),
		refundType, "11",
		create.Amount,
		create.PayInDate,
		create.CreatedAt)

	log.Println("Refund success")
	s, _ := json.Marshal(url)
	log.Println(string(s))
	return nil
}
