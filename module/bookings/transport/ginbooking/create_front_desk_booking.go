package ginbooking

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	bookingbiz "h5travelotobackend/module/bookings/biz"
	bookingmodel "h5travelotobackend/module/bookings/model"
	bookingsqlstorage "h5travelotobackend/module/bookings/storage"
	roomtypesqlstorage "h5travelotobackend/module/roomtypes/storage/sqlstorage"
	"net/http"
)

func CreateFrontDeskBooking(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var bookingCreate bookingmodel.FrontDeskBookingCreate

		if err := c.ShouldBind(&bookingCreate); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		bookingCreate.Booking.UnMask()
		bookingCreate.Booking.UserId = 0

		store := bookingsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())

		typeStore := roomtypesqlstorage.NewSqlStore(appCtx.GetGormDbConnection())

		biz := bookingbiz.NewCreateBookingBiz(store, typeStore, appCtx.GetPubSub())
		if err := biz.CreateFrontDeskBooking(c.Request.Context(), &bookingCreate); err != nil {
			panic(err)
		}

		bookingCreate.Booking.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(bookingCreate.Booking.FakeId))

	}
}
