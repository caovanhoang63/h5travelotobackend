package ginbooking

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	bookingbiz "h5travelotobackend/module/bookings/biz"
	bookingsqlstorage "h5travelotobackend/module/bookings/storage"
	"net/http"
)

type Filter struct {
	Date *common.CivilDate `json:"date" form:"date" binding:"required"`
}

func OverviewByDate(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var date Filter
		err = c.ShouldBind(&date)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := bookingsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := bookingbiz.NewStatisticBookingBiz(store)
		data, err := biz.OverviewByDate(c.Request.Context(), int(uid.GetLocalID()), date.Date)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}
