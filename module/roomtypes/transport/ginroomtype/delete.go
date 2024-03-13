package ginroomtype

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	roomtypebiz "h5travelotobackend/module/roomtypes/biz"
	roomtypesqlstorage "h5travelotobackend/module/roomtypes/storage/sqlstorage"
	"net/http"
)

func DeleteRoomType(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {

		roomUid, err := common.FromBase58(context.Param("room-type-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		roomId := int(roomUid.GetLocalID())

		store := roomtypesqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := roomtypebiz.NewDeleteRoomTypeBiz(store)
		if err := biz.DeleteRoomType(context.Request.Context(), roomId); err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
