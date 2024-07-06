package ginbookingdetail

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	bookingdetailbiz "h5travelotobackend/module/bookingdetails/biz"
	bookingdetailmodel "h5travelotobackend/module/bookingdetails/model"
	bookingdetailstorage "h5travelotobackend/module/bookingdetails/storage"
	"net/http"
)

func ListRoomsOfBooking(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("booking-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		bookingId := uid.GetLocalID()
		store := bookingdetailstorage.NewSqlStore(appCtx.GetGormDbConnection())

		biz := bookingdetailbiz.NewListRoomOfBookingBiz(store)

		var room []bookingdetailmodel.Room
		if room, err = biz.ListRoomOfBooking(c.Request.Context(), int(bookingId)); err != nil {
			panic(err)
		}
		for i := range room {
			room[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(room))

	}
}
