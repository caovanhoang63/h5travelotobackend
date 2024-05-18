package ginbooking

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	bookingbiz "h5travelotobackend/module/bookings/biz"
	bookingsqlstorage "h5travelotobackend/module/bookings/storage"
	dealsqlstorage "h5travelotobackend/module/deals/storage"
)

func UpdateDeal(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		bkUid, err := common.FromBase58(c.Param("booking-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		deadId, err := common.FromBase58(c.Param("deal-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := bookingsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		dealStore := dealsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := bookingbiz.NewAddDealBiz(store, dealStore)
		if err = biz.AddDeal(c.Request.Context(), int(bkUid.GetLocalID()), int(deadId.GetLocalID())); err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
