package skiochat

import (
	socketio "github.com/googollee/go-socket.io"
	"golang.org/x/net/context"
	chatbiz "h5travelotobackend/chat/module/room/biz"
	chatmessage "h5travelotobackend/chat/module/room/model/message"
	chatstorage "h5travelotobackend/chat/module/room/storage"
	"h5travelotobackend/common"
	"log"
)

func MessageSent(appCtx common.SimpleAppContext, rtEngine common.SimpleRealtimeEngine) func(s socketio.Conn, sent *chatmessage.MessageSent) {
	return func(s socketio.Conn, sent *chatmessage.MessageSent) {
		user := s.Context().(common.Requester)

		sent.Message.From = user.GetUserId()
		sent.Message.FromFakeId = common.NewUIDP(uint32(user.GetUserId()), common.DbTypeUser, 0)
		sent.Message.OnCreate()

		log.Printf("user %v sent: %s\n", user.GetUserId(), sent.Message.Message)
		log.Println("room id: ", sent.RoomId)

		store := chatstorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := chatbiz.NewCreateNewMessageBiz(store)

		if err := biz.CreateMessage(context.TODO(), sent.RoomId, sent.Message); err != nil {
			s.Emit(common.EventCannotSendMessage, true)
		}

		err := rtEngine.EmitToRoom(sent.RoomId, common.EventNewMessage, sent.Message)
		if err != nil {
			s.Emit(common.EventCannotSendMessage, true)
		}
	}
}
