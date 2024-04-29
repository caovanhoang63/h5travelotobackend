package ginchat

import (
	"github.com/gin-gonic/gin"
	chatbiz "h5travelotobackend/chat/module/room/biz"
	chatstorage "h5travelotobackend/chat/module/room/storage"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"net/http"
)

//url: /chat/hotels/:hotel_id

func GetChatRoom(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := chatstorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := chatbiz.NewFindChatRoomBiz(store)

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		room, err := biz.FindChatRoom(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}
		room.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(room))
	}
}
