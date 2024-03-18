package ginroomtypeabout

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	roomtypeaboutbiz "h5travelotobackend/module/roomtypeabout/biz"
	roomtypeaboutmongostorage "h5travelotobackend/module/roomtypeabout/storage/mongostorage"
)

func DeleteByRoomTypeId(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, roomTypeId, err := common.GetHotelAndRoomTypeId(c)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := roomtypeaboutmongostorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := roomtypeaboutbiz.NewDeleteRoomTypeAboutBit(store)
		if err := biz.DeleteRoomTypeAbout(c.Request.Context(), roomTypeId); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
