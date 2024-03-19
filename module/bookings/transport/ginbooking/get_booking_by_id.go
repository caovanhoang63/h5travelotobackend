package ginbooking

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	bookingbiz "h5travelotobackend/module/bookings/biz"
	bookingsqlstorage "h5travelotobackend/module/bookings/storage"
	"net/http"
)

func GetBookingById(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("booking-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := bookingsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := bookingbiz.NewFindBookingBiz(store)
		data, err := biz.GetBookingById(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		data.Mask(false)
		fmt.Println("hotel id ", data.HotelId)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
