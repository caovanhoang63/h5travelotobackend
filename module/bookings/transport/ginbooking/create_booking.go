package ginbooking

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	bookingbiz "h5travelotobackend/module/bookings/biz"
	"h5travelotobackend/module/bookings/bookingmodel"
	bookingsqlstorage "h5travelotobackend/module/bookings/storage"
	roomtypesqlstorage "h5travelotobackend/module/roomtypes/storage/sqlstorage"
	"net/http"
)

func CreateBooking(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bookingCreate bookingmodel.BookingCreate

		if err := c.ShouldBind(&bookingCreate); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		bookingCreate.UnMask()
		bookingCreate.UserId = requester.GetUserId()

		fmt.Println("hotel id ", bookingCreate.HotelId)
		store := bookingsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())

		// TODO: loai bo phu thuoc doi voi lop roomtype
		typeStore := roomtypesqlstorage.NewSqlStore(appCtx.GetGormDbConnection())

		biz := bookingbiz.NewCreateBookingBiz(store, typeStore, appCtx.GetPubSub())
		if err := biz.Create(c.Request.Context(), &bookingCreate, requester); err != nil {
			panic(err)
		}

		bookingCreate.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(bookingCreate.FakeId))

	}
}
