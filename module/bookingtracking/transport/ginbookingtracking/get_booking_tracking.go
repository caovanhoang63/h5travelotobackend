package ginbookingtracking

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	bookingtrackingbiz "h5travelotobackend/module/bookingtracking/biz"
	bookingtrackingstorage "h5travelotobackend/module/bookingtracking/storage"
	"net/http"
)

func GetStatesOfBooking(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookingUid, err := common.FromBase58(c.Param("booking-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := bookingtrackingstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := bookingtrackingbiz.NewGetBookingTrackingsBiz(store)
		data, err := biz.GetBookingTrackings(c.Request.Context(), int(bookingUid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
