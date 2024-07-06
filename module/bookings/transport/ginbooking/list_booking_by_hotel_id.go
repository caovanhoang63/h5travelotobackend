package ginbooking

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	bookingdetailstorage "h5travelotobackend/module/bookingdetails/storage"
	bookingbiz "h5travelotobackend/module/bookings/biz"
	"h5travelotobackend/module/bookings/model"
	bookingsqlstorage "h5travelotobackend/module/bookings/storage"
	"net/http"
)

func ListBookingHotelId(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.FullFill()

		var filter bookingmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter.HotelId = int(uid.GetLocalID())

		store := bookingsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		dbStore := bookingdetailstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := bookingbiz.NewListBookingBiz(store, dbStore)
		data, err := biz.ListBookingByHotelId(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))

	}
}
