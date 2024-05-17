package skio

import (
	"context"
	"errors"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"h5travelotobackend/chat/module/message/transport/skiochatmessage"
	"h5travelotobackend/chat/module/room/transport/skiochatroom"
	"h5travelotobackend/common"
	"h5travelotobackend/component/tokenprovider/jwt"
	userstorage "h5travelotobackend/module/users/storage"
)

func Setup(appCtx AppContext, engine *rtEngine) {
	server := engine.server

	// handles when a client open a connection to server
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected: ", s.ID(), "Ip: ", s.RemoteAddr(), s.ID())
		return nil
	})

	// handles when the connection has error
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	// handles when connection close
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	server.OnEvent("/", "ping", func(s socketio.Conn) {
		s.Emit("pong")
	})

	// authenticate messages
	// 1. Check the access token
	// 	1.1. If the access token or userId is invalid, sends message
	// 		 "authentication_failed" to client and close the connection
	//  1.2. If having no trouble, create and save a new AppSocket
	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
		db := appCtx.GetGormDbConnection()
		store := userstorage.NewSqlStore(db)
		tokenProvider := jwt.NewJWTProvider(appCtx.GetSecretKey())

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		if user.Status == 0 {
			s.Emit("authentication_failed", errors.New("you has been banned/deleted"))
			s.Close()
			return
		}
		user.Mask(false)

		appSck := NewAppSocket(s, user)
		s.SetContext(user)
		engine.saveAppSocket(user.Id, appSck)
		fmt.Println("socket authenticated")

		s.Emit("authenticated", user)

	})

	server.OnEvent("/", common.EventMessageSent, skiochatmessage.MessageSent(appCtx, engine))
	server.OnEvent("/", common.EventUserJoined, skiochatroom.UserJoined(appCtx, engine))
	server.OnEvent("/", common.EventCustomerSeenMessage, skiochatroom.UserSeenMessages(appCtx, engine))
	server.OnEvent("/", common.EventHotelSeenMessage, skiochatroom.HotelSeenMessages(appCtx, engine))

}
