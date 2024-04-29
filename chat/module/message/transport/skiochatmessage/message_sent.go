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
) func(s socketio.Conn, sent *chatmessagemodel.MessageSent) {
	return func(s socketio.Conn, sent *chatmessagemodel.MessageSent) {
		user := s.Context().(common.Requester)

		sent.Message.UserId = user.GetUserId()
		sent.Message.UserFakeId = common.NewUIDP(uint32(user.GetUserId()), common.DbTypeUser, 0)
		sent.Message.OnCreate()

		log.Printf("user %v sent: %s\n", user.GetUserId(), sent.Message.Message)
		log.Println("room id: ", sent.RoomId)

		store := chatmessagestorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := chatmessagebiz.NewCreateNewMessageBiz(store)

		if err := biz.CreateMessage(context.TODO(), sent.RoomId, sent.Message); err != nil {
			s.Emit(common.EventCannotSendMessage, true)
		}

		err := rtEngine.EmitToRoom(sent.RoomId, common.EventNewMessage, sent.Message)
		if err != nil {
			s.Emit(common.EventCannotSendMessage, true)
		}
	}
}
