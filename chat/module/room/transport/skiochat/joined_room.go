package skiochat

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"h5travelotobackend/common"
)

func UserJoined(appCtx common.SimpleAppContext,
	rtEngine common.SimpleRealtimeEngine) func(s socketio.Conn, roomId string) {
	return func(s socketio.Conn, roomId string) {
		fmt.Println("User joined room", roomId)
		s.Join(roomId)
		s.Emit("joined", roomId)
	}
}
