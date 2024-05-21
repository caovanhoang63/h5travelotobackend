package ginpayin

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/payment/vnpay"
	"h5travelotobackend/module/bookings/transport/bklocalhandler"
	payinbiz "h5travelotobackend/payment/module/payin/biz"
	payinstore "h5travelotobackend/payment/module/payin/store"
	"h5travelotobackend/payment/module/paymentevent/transport/pelocalhandler"
	"log"
	"net/http"
)

func VnpIPN(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		vnp := appCtx.GetVnPay()
		var ipn vnpay.IPNRequest
		err := c.ShouldBind(&ipn)
		if err != nil {
			c.JSON(http.StatusBadRequest, vnpay.NewOtherError())
			return
		}

		if !vnp.CheckSum(c.Request.URL.String()) {
			c.JSON(http.StatusBadRequest, vnpay.NewInvalidCheckSum())
			return
		}
		store := payinstore.NewStore(appCtx.GetGormDbConnection())
		peStore := pelocalhandler.NewPELocalHandler(appCtx)
		bkStore := bklocalhandler.NewCountBookedRoomLocalHandler(appCtx)
		successBiz := payinbiz.NewSuccessPBBiz(store, appCtx.GetPubSub())
		biz := payinbiz.NewVnpIPNBiz(store, peStore, bkStore, appCtx.GetVnPay(), successBiz)

		log.Println(c.Request.URL.String())
		c.JSON(http.StatusOK, *biz.HandleIPNRequest(c.Request.Context(), &ipn))

	}
}
