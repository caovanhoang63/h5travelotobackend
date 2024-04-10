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

func ConfirmBookingTracking(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "confirm booking when select rooms ",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var booking common.DTORoomType
			err := json.Unmarshal(message.Data, &booking)
			if err != nil {
				return err
			}
			data := bookingtrackingmodel.BookingTrackingCreate{
				BookingId: booking.Id,
				State:     "confirmed",
			}
			store := bookingtrackingstorage.NewSqlStore(appCtx.GetGormDbConnection())
			return store.Create(ctx, &data)
		},
	}
}
