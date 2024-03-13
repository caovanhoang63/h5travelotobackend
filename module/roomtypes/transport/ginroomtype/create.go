package ginroomtype

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	roomtypebiz "h5travelotobackend/module/roomtypes/biz"
	roomtypemodel "h5travelotobackend/module/roomtypes/model"
	roomtypesqlstorage "h5travelotobackend/module/roomtypes/storage/sqlstorage"
	"net/http"
)

func CreateRoomType(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		hotelUid, err := common.FromBase58(context.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		hotelId := int(hotelUid.GetLocalID())

		var data roomtypemodel.RoomTypeCreate

		if err := context.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		data.HotelId = hotelId
		store := roomtypesqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := roomtypebiz.NewRoomTypeBiz(store)
		if err := biz.CreateRoomType(context.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))
	}
}
