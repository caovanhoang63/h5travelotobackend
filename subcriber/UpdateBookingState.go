package subcriber

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	bookingtrackingmodel "h5travelotobackend/module/bookingtracking/model"
	bookingtrackingstorage "h5travelotobackend/module/bookingtracking/storage"
)

func TrackingBookingStateUpdate(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Tracking Booking State Update",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var tracking bookingtrackingmodel.BookingTrackingCreate
			err := json.Unmarshal(message.Data, &tracking)
			tracking.UnMask()
			if err != nil {
				return err
			}
			store := bookingtrackingstorage.NewSqlStore(appCtx.GetGormDbConnection())
			return store.Create(ctx, &tracking)
		},
	}
}
