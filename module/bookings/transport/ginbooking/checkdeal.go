package ginbooking

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	bookingbiz "h5travelotobackend/module/bookings/biz"
	bookingmodel "h5travelotobackend/module/bookings/model"
	dealsqlstorage "h5travelotobackend/module/deals/storage"
	roomtypesqlstorage "h5travelotobackend/module/roomtypes/storage/sqlstorage"
)

// bookings/checkdeal/:deal-id

func CheckDeal(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("deal-id"))

		var create bookingmodel.BookingCreate
		if err := c.ShouldBind(&create); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		create.UnMask()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		dealStore := dealsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		typeStore := roomtypesqlstorage.NewSqlStore(appCtx.GetGormDbConnection())

		biz := bookingbiz.NewCheckDealBiz(dealStore, typeStore)

		if err = biz.CheckDeal(c.Request.Context(), &create, int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(create))
	}
}
