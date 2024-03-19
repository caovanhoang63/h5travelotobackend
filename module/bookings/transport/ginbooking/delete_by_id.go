package ginbooking

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	bookingbiz "h5travelotobackend/module/bookings/biz"
	bookingsqlstorage "h5travelotobackend/module/bookings/storage"
	workersqlstorage "h5travelotobackend/module/worker/storage/sqlstorage"
	"net/http"
)

func DeleteBookingById(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("booking-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := bookingsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		workerStore := workersqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := bookingbiz.NewDeleteBookingBiz(store, store, workerStore)
		if err := biz.DeleteBooking(c.Request.Context(), requester, int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
