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

func ListBookingByUserId(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("user-id"))
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

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := bookingsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		dbStore := bookingdetailstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := bookingbiz.NewListBookingBiz(store, dbStore)
		data, err := biz.ListBookingByUserId(c.Request.Context(), requester, int(uid.GetLocalID()), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))

	}
}
