package ginroom

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	roombiz "h5travelotobackend/module/rooms/biz"
	roommodel "h5travelotobackend/module/rooms/model"
	roomstorage "h5travelotobackend/module/rooms/storage"
	roomtypesqlstorage "h5travelotobackend/module/roomtypes/storage/sqlstorage"
	"net/http"
)

func CreateRoom(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data roommodel.RoomCreate
		if err := context.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Check hotel id
		hotelUid, err := common.FromBase58(context.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		data.HotelId = int(hotelUid.GetLocalID())

		// Check room type id
		roomTypeUid, err := common.FromBase58(data.RoomTypeFakeId)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		data.RoomTypeID = int(roomTypeUid.GetLocalID())

		store := roomstorage.NewSqlStore(appCtx.GetGormDbConnection())
		//TODO: use gRPC instead of SQL storage
		findRoomTypeStore := roomtypesqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := roombiz.NewCreateRoomBiz(store, findRoomTypeStore, appCtx.GetPubSub())

		if err := biz.CreateRoom(context.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))
	}
}
