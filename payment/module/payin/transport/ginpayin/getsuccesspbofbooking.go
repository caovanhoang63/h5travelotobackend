package ginpayin

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	payinbiz "h5travelotobackend/payment/module/payin/biz"
	payinstore "h5travelotobackend/payment/module/payin/store"
	"net/http"
)

func GetPaymentSuccessBooking(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("booking_id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		store := payinstore.NewStore(appCtx.GetGormDbConnection())
		biz := payinbiz.NewFindPaymentBookingBiz(store)
		pb, err := biz.FindPBSuccessOfBooking(c.Request.Context(), requester, int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}
		pb.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(pb))

	}
}
