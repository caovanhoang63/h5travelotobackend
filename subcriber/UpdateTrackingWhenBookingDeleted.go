package subcriber

import (
	"context"
	"encoding/json"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	bookingtrackingmodel "h5travelotobackend/module/bookingtracking/model"
	bookingtrackingstorage "h5travelotobackend/module/bookingtracking/storage"
)

func DeleteTrackingWhenBookingDeleted(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Update tracking state to deleted when a booking is deleted",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var booking common.DTORoomType
			err := json.Unmarshal(message.Data, &booking)
			if err != nil {
				return err
			}

			store := bookingtrackingstorage.NewSqlStore(appCtx.GetGormDbConnection())
			return store.Update(ctx, booking.Id, &bookingtrackingmodel.BookingTrackingUpdate{
				State: "deleted",
			})
		},
	}
}
