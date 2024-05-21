package payinbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	payinmodel "h5travelotobackend/payment/module/payin/model"
	"log"
)

type successPBStatusBiz struct {
	pbStore UpdatPBStatusStore
	pb      pubsub.Pubsub
}

func NewSuccessPBBiz(pbStore UpdatPBStatusStore, pb pubsub.Pubsub) *successPBStatusBiz {
	return &successPBStatusBiz{pbStore: pbStore, pb: pb}
}

func (biz *successPBStatusBiz) SuccessStatusBiz(ctx context.Context, update *payinmodel.PaymentBooking) error {
	err := biz.pbStore.UpdateStatus(ctx, update.TxnId,
		&payinmodel.PaymentBookingUpdateStatus{PaymentStatus: common.PaymentStatusSuccess})
	if err != nil {
		return common.ErrInternal(err)
	}
	message := pubsub.NewMessage(common.PaymentBooking{
		TxnId:     update.TxnId,
		BookingId: update.BookingId,
		Amount:    update.Amount,
	})
	err = biz.pb.Publish(ctx, common.TopicPaymentSuccess, message)
	if err != nil {
		log.Println("Error publish message", err)
	}
	return nil
}
