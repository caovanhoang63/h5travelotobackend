package ginroom

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	bookingdetailrepo "h5travelotobackend/module/bookingdetails/repo"
	bookingdetailstorage "h5travelotobackend/module/bookingdetails/storage"
	bookingsqlstorage "h5travelotobackend/module/bookings/storage"
	bookingdetailbiz "h5travelotobackend/module/rooms/biz"
	roomstorage "h5travelotobackend/module/rooms/storage"
	"net/http"
	"time"
)

// GetAvailableRoomForBooking url
// {{}}/v1/hotels/:hotel-id/available-room?start-date=2021-07-01&&end-date=2021-07-02&room-type-id=1

const layout = "2006-01-02"

func GetAvailableRoomByDate(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		StartDate, err := time.Parse(layout, c.Query("start-date"))
		EndDate, err := time.Parse(layout, c.Query("end-date"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		uid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		hotelId := int(uid.GetLocalID())

		roomStore := roomstorage.NewSqlStore(appCtx.GetGormDbConnection())
		store := bookingdetailstorage.NewSqlStore(appCtx.GetGormDbConnection())
		bookingStore := bookingsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		repo := bookingdetailrepo.NewGetRoomBookedRepo(store)
		biz := bookingdetailbiz.NewListAvailableRoom(repo, roomStore, bookingStore)

		data, err := biz.ListAvailableRoomByDate(c.Request.Context(), &StartDate, &EndDate, hotelId)
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}
