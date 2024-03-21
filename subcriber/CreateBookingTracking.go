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

func CreateBookingTracking(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Create new booking tracking when a booking is created",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var booking common.DTORoomType
			err := json.Unmarshal(message.Data, &booking)
			if err != nil {
				return err
			}
			data := bookingtrackingmodel.BookingTrackingCreate{
				SqlModel:      common.SqlModel{},
				BookingId:     booking.Id,
				BookingFakeId: common.UID{},
				State:         "pending",
			}
			store := bookingtrackingstorage.NewSqlStore(appCtx.GetGormDbConnection())
			return store.Create(ctx, &data)
		},
	}
}
