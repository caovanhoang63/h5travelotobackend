package subcriber

import (
	"context"
	"encoding/json"
	"fmt"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	roomtypesqlstorage "h5travelotobackend/module/roomtypes/storage/sqlstorage"
)

func DecreaseTotalRoomWhenCreateNewRoom(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Decrease total room when delete room",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var roomType common.DTORoomType
			if err := json.Unmarshal(message.Data, &roomType); err != nil {
				return err
			}
			db := appCtx.GetGormDbConnection()
			store := roomtypesqlstorage.NewSqlStore(db)
			fmt.Println("roomType.Id", roomType.Id)
			return store.DecreaseTotalRoom(ctx, roomType.Id)
		},
	}
}
