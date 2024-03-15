package ginroom

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	roombiz "h5travelotobackend/module/rooms/biz"
	roommodel "h5travelotobackend/module/rooms/model"
	roomstorage "h5travelotobackend/module/rooms/storage"
	"net/http"
	"strconv"
)

func ListRoomWithCondition(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var filter roommodel.Filter
		var paging common.Paging

		hotelUid, err := common.FromBase58(context.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if str, ok := context.Params.Get("room-type-id"); ok {
			roomTypeUid, err := common.FromBase58(str)
			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}
			filter.RoomTypeId = int(roomTypeUid.GetLocalID())
		}

		if str, ok := context.Params.Get("status"); ok {
			filter.Status, err = strconv.Atoi(str)
			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}
		} else {
			filter.Status = 1
		}

		filter.HotelId = int(hotelUid.GetLocalID())

		if err := context.ShouldBindQuery(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.FullFill()

		store := roomstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := roombiz.NewListRoomBiz(store)
		data, err := biz.ListRoomWithCondition(context.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
