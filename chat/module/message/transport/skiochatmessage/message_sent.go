package skiochatmessage

import (
	socketio "github.com/googollee/go-socket.io"
	"golang.org/x/net/context"
	chatmessagebiz "h5travelotobackend/chat/module/message/biz"
	chatmessagemodel "h5travelotobackend/chat/module/message/model"
	chatmessagestorage "h5travelotobackend/chat/module/message/storage"
	"h5travelotobackend/common"
)

func MessageSent(appCtx common.SimpleAppContext,
	rtEngine common.SimpleRealtimeEngine,
) func(s socketio.Conn, message *chatmessagemodel.Message) {
	return func(s socketio.Conn, message *chatmessagemodel.Message) {
		user := s.Context().(common.Requester)
		message.UserId = user.GetUserId()
		if user.GetRole() == common.RoleCustomer {
			message.IsFromCustomer = true
		} else {
			message.IsFromCustomer = false
		}
		message.UserFakeId = common.NewUIDP(uint32(user.GetUserId()), common.DbTypeUser, 0)
		store := chatmessagestorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := chatmessagebiz.NewCreateNewMessageBiz(store, appCtx.GetPubSub())
		if err := biz.CreateMessage(context.TODO(), message); err != nil {
			s.Emit(common.EventCannotSendMessage, true)
		}
		roomId, err := message.RoomId.MarshalText()
		if err != nil {
			s.Emit(common.EventCannotSendMessage, true)
		}
		err = rtEngine.EmitToRoom(string(roomId), common.EventNewMessage, message)
		if err != nil {
			s.Emit(common.EventCannotSendMessage, true)
		}
	}
}
