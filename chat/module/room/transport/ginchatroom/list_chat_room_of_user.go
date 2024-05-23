package ginchatroom

import (
	"github.com/gin-gonic/gin"
	chatbiz "h5travelotobackend/chat/module/room/biz"
	chatstorage "h5travelotobackend/chat/module/room/storage"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelstorage "h5travelotobackend/search/module/hotel/storage/esstore"
	"net/http"
)

func ListChatRoomByUser(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.FullFill()
		store := chatstorage.NewMongoStore(appCtx.GetMongoConnection())
		hStore := hotelstorage.NewESStore(appCtx.GetElasticSearchClient())
		biz := chatbiz.NewListChatRoomByUserBiz(store, hStore)
		data, err := biz.ListChatRoomByUser(c.Request.Context(), requester, &paging)
		if err != nil {
			panic(err)
		}
		for i := range data {
			data[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}
