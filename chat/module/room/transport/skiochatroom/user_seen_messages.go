package skiochatroom

import (
	socketio "github.com/googollee/go-socket.io"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	chatroombiz "h5travelotobackend/chat/module/room/biz"
	chatroomstorage "h5travelotobackend/chat/module/room/storage"
	"h5travelotobackend/common"
)

func UserSeenMessages(appCtx common.SimpleAppContext,
	rtEngine common.SimpleRealtimeEngine) func(s socketio.Conn, roomId string) {
	return func(s socketio.Conn, roomId string) {

		id, err := primitive.ObjectIDFromHex(roomId)
		if err != nil {
			s.Emit(common.EventCannotMarkMessageAsSeen, roomId)
		}

		store := chatroomstorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := chatroombiz.NewCustomerSeenMessageBiz(store)

		if err := biz.CustomerSeenMessages(context.Background(), &id); err != nil {
			s.Emit(common.EventCannotMarkMessageAsSeen, roomId)
		}

		err = rtEngine.EmitToRoom(roomId, common.EventCustomerSeenMessage, roomId)
		if err != nil {
			return
		}
	}
}
