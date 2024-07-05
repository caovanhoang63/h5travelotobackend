package subcriber

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	bookingsqlstorage "h5travelotobackend/module/bookings/storage"
)

func UpdateBookingStateToPaid(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Update booking state to paid when payment booking done",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var data common.PaymentBooking
			err := json.Unmarshal(message.Data, &data)
			if err != nil {
				return err
			}
			store := bookingsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
			return store.UpdateStateToPaid(ctx, data.BookingId)
		},
	}
}
