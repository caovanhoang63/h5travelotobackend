package skiochatmessage

import (
	socketio "github.com/googollee/go-socket.io"
	"golang.org/x/net/context"
	chatmessagebiz "h5travelotobackend/chat/module/message/biz"
	chatmessagemodel "h5travelotobackend/chat/module/message/model"
	chatmessagestorage "h5travelotobackend/chat/module/message/storage"
	"h5travelotobackend/common"
	"log"
)

func MessageSent(appCtx common.SimpleAppContext,
	rtEngine common.SimpleRealtimeEngine,
) func(s socketio.Conn, message *chatmessagemodel.Message) {
	return func(s socketio.Conn, message *chatmessagemodel.Message) {
		user := s.Context().(common.Requester)
		message.UserId = user.GetUserId()
		message.UserFakeId = common.NewUIDP(uint32(user.GetUserId()), common.DbTypeUser, 0)

		log.Printf("user %v sent: %s\n", user.GetUserId(), message.Message)
		log.Println("room id: ", message.RoomId)

		store := chatmessagestorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := chatmessagebiz.NewCreateNewMessageBiz(store, appCtx.GetPubSub())

		if err := biz.CreateMessage(context.TODO(), message); err != nil {
			s.Emit(common.EventCannotSendMessage, true)
		}

		err := rtEngine.EmitToRoom(message.RoomId.String(), common.EventNewMessage, message.Message)
		if err != nil {
			s.Emit(common.EventCannotSendMessage, true)
		}
	}
}
