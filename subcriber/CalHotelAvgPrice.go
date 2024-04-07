package subcriber

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	hotelstorage "h5travelotobackend/module/hotels/storage"
	"log"
)

func CalAvgHotelPrice(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "update total room type and avg price of hotel when create new room type",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var Room common.DTORoomType
			err := json.Unmarshal(message.Data, &Room)
			if err != nil {
				log.Println(err)
			}
			store := hotelstorage.NewSqlStore(appCtx.GetGormDbConnection())

			return store.UpdateTotalRoomTypeAndAvgPrice(ctx, &Room)
		},
	}
}
