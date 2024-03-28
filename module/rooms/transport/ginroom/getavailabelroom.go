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
)

// GetAvailableRoom url
// {{}}/v1/hotels/:hotel-id/booking/:booking-id/available-room

func GetAvailableRoom(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("booking-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		roomStore := roomstorage.NewSqlStore(appCtx.GetGormDbConnection())
		store := bookingdetailstorage.NewSqlStore(appCtx.GetGormDbConnection())
		bookingStore := bookingsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		repo := bookingdetailrepo.NewGetRoomBookedRepo(store)
		biz := bookingdetailbiz.NewListAvailableRoom(repo, roomStore, bookingStore)

		data, err := biz.ListAvailableRoom(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
