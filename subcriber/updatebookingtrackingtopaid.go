package subcriber

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	bookingtrackingmodel "h5travelotobackend/module/bookingtracking/model"
	bookingtrackingstorage "h5travelotobackend/module/bookingtracking/storage"
)

func UpdateBookingTrackingToPaid(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Insert booking state paid when payment booking done",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var data common.PaymentBooking
			err := json.Unmarshal(message.Data, &data)
			if err != nil {
				return err
			}
			store := bookingtrackingstorage.NewSqlStore(appCtx.GetGormDbConnection())
			return store.Create(ctx, &bookingtrackingmodel.BookingTrackingCreate{
				BookingId: data.BookingId,
				State:     common.BookingStatePaid,
			})
		},
	}
}
