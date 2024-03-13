package ginroom

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	roombiz "h5travelotobackend/module/rooms/biz"
	roommodel "h5travelotobackend/module/rooms/model"
	roomstorage "h5travelotobackend/module/rooms/storage"
	"net/http"
)

func UpdateRoom(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data roommodel.RoomUpdate
		roomUid, err := common.FromBase58(context.Param("room-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := context.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := roomstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := roombiz.NewUpdateRoomBiz(store)
		if err := biz.UpdateRoom(context.Request.Context(), int(roomUid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
