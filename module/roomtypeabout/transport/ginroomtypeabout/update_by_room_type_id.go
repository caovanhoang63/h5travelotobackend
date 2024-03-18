package ginroomtypeabout

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	roomtypeaboutbiz "h5travelotobackend/module/roomtypeabout/biz"
	roomtypeaboutmodel "h5travelotobackend/module/roomtypeabout/model"
	roomtypeaboutmongostorage "h5travelotobackend/module/roomtypeabout/storage/mongostorage"
	"net/http"
)

func UpdateByRoomTypeId(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, roomTypeId, err := common.GetHotelAndRoomTypeId(c)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var updateDate roomtypeaboutmodel.RoomTypeAboutUpdate
		if err := c.ShouldBind(&updateDate); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := roomtypeaboutmongostorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := roomtypeaboutbiz.NewUpdateRoomTypeAboutBiz(store)
		if err := biz.UpdateRoomTypeAbout(c.Request.Context(), roomTypeId, &updateDate); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
