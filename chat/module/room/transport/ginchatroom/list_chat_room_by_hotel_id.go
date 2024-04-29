package ginchatroom

import (
	"github.com/gin-gonic/gin"
	chatbiz "h5travelotobackend/chat/module/room/biz"
	chatstorage "h5travelotobackend/chat/module/room/storage"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"net/http"
)

func ListChatRoomByHotelId(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("hotel-id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.FullFill()

		hotelId := int(uid.GetLocalID())

		store := chatstorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := chatbiz.NewListChatRoomBiz(store)

		data, err := biz.ListChatRoomByHotelId(c.Request.Context(), hotelId, &paging)

		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}
