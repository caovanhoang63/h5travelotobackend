package ginchatmessage

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	chatmessagebiz "h5travelotobackend/chat/module/message/biz"
	chatmessagestorage "h5travelotobackend/chat/module/message/storage"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"net/http"
)

// url : /chat/rooms/:room-id/messages

func ListMessagesByRoomId(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomId := c.Param("room-id")
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.FullFill()
		store := chatmessagestorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := chatmessagebiz.NewListMessageBiz(store)

		id, err := primitive.ObjectIDFromHex(roomId)
		messages, err := biz.ListMessageByRoomId(c.Request.Context(), &id, &paging)
		if err != nil {
			panic(err)
		}
		for i := range messages {
			messages[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(messages, paging, nil))
	}
}
