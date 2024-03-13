package ginroom

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	roombiz "h5travelotobackend/module/rooms/biz"
	roomstorage "h5travelotobackend/module/rooms/storage"
	"net/http"
)

func DeleteRoom(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(err)
		}

		store := roomstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := roombiz.NewDeleteRoomBiz(store)
		if err := biz.DeleteRoom(context.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
