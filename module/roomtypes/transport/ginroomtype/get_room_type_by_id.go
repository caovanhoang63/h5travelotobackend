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

func GetRoomTypeById(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data *roomtypemodel.RoomType

		sqlStore := roomtypesqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := roomtypebiz.NewFindRoomTypeBiz(sqlStore)
		data, err = biz.GetRoomTypeByID(context.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		data.Mask(false)
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
