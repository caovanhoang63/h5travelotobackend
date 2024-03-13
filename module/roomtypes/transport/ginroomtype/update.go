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

func UpdateRoomType(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		uid, err := common.FromBase58(context.Param("room-type-id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data roomtypemodel.RoomTypeUpdate

		if err := context.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := roomtypesqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := roomtypebiz.NewRoomTypeUpdateBiz(store)
		if err := biz.Update(context.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
