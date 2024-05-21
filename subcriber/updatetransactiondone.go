package subcriber

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	pestore "h5travelotobackend/payment/module/paymentevent/store"
)

func UpdateTransactionDone(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Update transaction done when payment booking done",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var data common.PaymentBooking
			err := json.Unmarshal(message.Data, &data)
			if err != nil {
				return err
			}
			store := pestore.NewStore(appCtx.GetGormDbConnection())
			return store.UpdateStatusToSuccess(ctx, data.TxnId)
		},
	}
}
