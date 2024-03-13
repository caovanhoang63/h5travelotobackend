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

func CreateRoom(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data roommodel.RoomCreate

		if err := context.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := roomstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := roombiz.NewCreateRoomBiz(store)
		if err := biz.CreateRoom(context.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))
	}
}
