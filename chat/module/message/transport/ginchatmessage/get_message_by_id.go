package ginchatmessage

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	chatmessagebiz "h5travelotobackend/chat/module/message/biz"
	chatmessagestorage "h5travelotobackend/chat/module/message/storage"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"log"
	"net/http"
)

// url: /chat/messages/:message-id

func GetMessageById(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		messageId := c.Param("message-id")

		store := chatmessagestorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := chatmessagebiz.NewFindMessageBiz(store)
		id, err := primitive.ObjectIDFromHex(messageId)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		message, err := biz.GetMessageById(c.Request.Context(), &id, c.MustGet(common.CurrentUser).(common.Requester))
		if err != nil {
			panic(err)
		}
		message.Mask(false)

		log.Println("message: ", message.IsFromCustomer)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(message))
	}
}
