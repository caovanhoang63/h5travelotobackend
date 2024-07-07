package ginbooking

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	bookingbiz "h5travelotobackend/module/bookings/biz"
	bookingsqlstorage "h5travelotobackend/module/bookings/storage"
	"net/http"
)

func CancelBooking(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("booking-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		bookingId := int(uid.GetLocalID())
		store := bookingsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := bookingbiz.NewUpdateStateBiz(store, appCtx.GetPubSub())

		err = biz.UpdateState(c.Request.Context(), bookingId, common.BookingStateCanceled)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
