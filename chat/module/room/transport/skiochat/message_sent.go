package skiochat

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	chatmessage "h5travelotobackend/chat/module/room/model/message"
	"h5travelotobackend/common"
)

func MessageSent(appCtx common.SimpleAppContext, rtEngine common.SimpleRealtimeEngine) func(s socketio.Conn, sent *chatmessage.MessageSent) {
	return func(s socketio.Conn, sent *chatmessage.MessageSent) {
		user := s.Context().(common.Requester)
		sent.Message.From = user.GetUserId()
		sent.Message.FromFakeId = common.NewUIDP(uint32(user.GetUserId()), common.DbTypeUser, 0)
		sent.Message.OnCreate()

		fmt.Printf("user %v sent: %s\n", user.GetUserId(), sent.Message.Message)
		fmt.Println("room id: ", sent.RoomId)
		err := rtEngine.EmitToRoom(sent.RoomId, common.EventNewMessage, sent.Message)
		if err != nil {
			s.Emit(common.EventCannotSendMessage, true)
		}
	}
}
