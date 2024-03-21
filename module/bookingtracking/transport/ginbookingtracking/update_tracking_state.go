package ginbookingtracking

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	bookingtrackingbiz "h5travelotobackend/module/bookingtracking/biz"
	bookingtrackingmodel "h5travelotobackend/module/bookingtracking/model"
	bookingtrackingstorage "h5travelotobackend/module/bookingtracking/storage"
	"net/http"
)

func UpdateTrackingState(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookingUid, err := common.FromBase58(c.Param("booking-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data bookingtrackingmodel.BookingTrackingUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := bookingtrackingstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := bookingtrackingbiz.NewUpdateTrackingBiz(store)
		if err := biz.UpdateTracking(c.Request.Context(), int(bookingUid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
