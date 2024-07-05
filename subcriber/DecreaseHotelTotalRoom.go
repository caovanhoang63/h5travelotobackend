package subcriber

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	hotelstorage "h5travelotobackend/module/hotels/storage"
)

func DecreaseHotelTotalRoomWhenCreateNewRoom(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Decrease hotel total room when create new room",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var roomType common.DTORoomType
			if err := json.Unmarshal(message.Data, &roomType); err != nil {
				return err
			}
			db := appCtx.GetGormDbConnection()
			store := hotelstorage.NewSqlStore(db)
			return store.DecreaseTotalRoom(ctx, roomType.HotelId)
		},
	}
}
