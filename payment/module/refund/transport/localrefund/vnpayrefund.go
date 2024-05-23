package localrefund

import (
	"h5travelotobackend/component/appContext"
)

type VnPayRefund struct {
	appCtx appContext.AppContext
}

//func (v *VnPayRefund) VnPayRefund(ctx context.Context, reason string, pb *payinmodel.PaymentBooking) {
//	create := &refundmodel.RefundBookingCreate{
//		BookingId:  pb.BookingId,
//		Reason:     reason,
//		PayInTxnId: pb.TxnId,
//		Currency:   pb.Currency,
//	}
//
//	vnp := v.appCtx.GetVnPay()
//
//	hStore := localhoteldetail.NewHotelDetailLocalHandler(v.appCtx)
//	rStore := refundstore.NewStore(v.appCtx.GetGormDbConnection())
//	peStore := paymenteventstore.NewStore(v.appCtx.GetGormDbConnection())
//
//	biz := refundbiz.NewRefundBiz(hStore, rStore, peStore)
//
//	err := biz.Refund(ctx)
//	if err != nil {
//		return
//	}
//
//}
