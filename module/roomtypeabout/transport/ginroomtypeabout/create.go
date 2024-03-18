package ginroomtypeabout

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	roomtypeaboutbiz "h5travelotobackend/module/roomtypeabout/biz"
	roomtypeaboutmodel "h5travelotobackend/module/roomtypeabout/model"
	roomtypeaboutmongostorage "h5travelotobackend/module/roomtypeabout/storage/mongostorage"
	roomtypesqlstorage "h5travelotobackend/module/roomtypes/storage/sqlstorage"
)

func CreateRoomTypeAbout(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, roomTypeId, err := common.GetHotelAndRoomTypeId(c)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data roomtypeaboutmodel.RoomTypeAbout
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		data.RoomTypeId = roomTypeId

		store := roomtypeaboutmongostorage.NewMongoStore(appCtx.GetMongoConnection())
		findStore := roomtypesqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := roomtypeaboutbiz.NewCreateRoomTypeAboutBit(store, findStore)
		if err := biz.CreateRoomTypeAbout(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
