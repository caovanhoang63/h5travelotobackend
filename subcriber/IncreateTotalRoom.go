package subcriber

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	roomtypesqlstorage "h5travelotobackend/module/roomtypes/storage/sqlstorage"
)

func IncreaseTotalRoomWhenCreateNewRoom(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Increase total room when create new room",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var roomType common.DTORoomType
			if err := json.Unmarshal(message.Data, &roomType); err != nil {
				return err
			}
			db := appCtx.GetGormDbConnection()
			store := roomtypesqlstorage.NewSqlStore(db)
			fmt.Println("roomType.Id", roomType.Id)
			return store.IncreaseTotalRoom(ctx, roomType.Id)
		},
	}
}
