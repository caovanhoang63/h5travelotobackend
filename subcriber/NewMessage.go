package subcriber

import (
	"encoding/json"
	"golang.org/x/net/context"
	chatroommodel "h5travelotobackend/chat/module/room/model"
	chatroomstorage "h5travelotobackend/chat/module/room/storage"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
)

func NewMessage(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Handle when new message is created",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var chatMessage chatroommodel.ChatMessage
			if err := json.Unmarshal(message.Data, &chatMessage); err != nil {
				return err
			}
			store := chatroomstorage.NewMongoStore(appCtx.GetMongoConnection())
			return store.HandleNewMessage(ctx, &chatMessage)
		},
	}
}
