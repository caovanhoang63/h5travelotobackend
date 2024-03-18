package ginroomtypeabout

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	roomtypeaboutbiz "h5travelotobackend/module/roomtypeabout/biz"
	roomtypeaboutmongostorage "h5travelotobackend/module/roomtypeabout/storage/mongostorage"
	"net/http"
)

func GetAboutByRoomTypeId(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, roomTypeId, err := common.GetHotelAndRoomTypeId(c)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		mStore := roomtypeaboutmongostorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := roomtypeaboutbiz.NewFindRoomTypeAboutBiz(mStore)
		data, err := biz.GetByRoomTypeId(c.Request.Context(), roomTypeId)

		if err != nil {
			panic(err)
		}

		data.FakeRoomTypeId = c.Param("room_type_id")

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
