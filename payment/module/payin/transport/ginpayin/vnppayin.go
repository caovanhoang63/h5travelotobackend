package ginpayin

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/module/bookings/transport/bklocalhandler"
	payinbiz "h5travelotobackend/payment/module/payin/biz"
	paymentmodel "h5travelotobackend/payment/module/payin/model"
	payinstore "h5travelotobackend/payment/module/payin/store"
	"h5travelotobackend/payment/module/paymentevent/transport/pelocalhandler"
	"log"
	"net/http"
)

func PayIn(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.Request.URL.String())
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		var info paymentmodel.PaymentBookingCreate
		if err := c.ShouldBind(&info); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := info.UnMask(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		info.Method = common.PaymentMethodVnPay
		store := payinstore.NewStore(appCtx.GetGormDbConnection())
		peStore := pelocalhandler.NewPELocalHandler(appCtx)
		bkStore := bklocalhandler.NewCountBookedRoomLocalHandler(appCtx)
		biz := payinbiz.NewPayInBiz(store, peStore, bkStore)
		err := biz.NewPaymentBooking(c, requester, &info)

		if err != nil {
			panic(err)
		}

		vnPay := appCtx.GetVnPay()
		url := vnPay.NewPayInUrl(info.Amount, info.Currency, info.BookingFakeId.String(), c.ClientIP(), info.TxnId, info.CreatedAt)

		response := paymentmodel.PaymentInfoResponse{
			PaymentUrl: url,
			Amount:     info.Amount,
			Currency:   info.Currency,
			BookingId:  info.BookingFakeId,
			Method:     common.PaymentMethodVnPay,
			TxnId:      info.TxnId,
			DealId:     info.DealFakeId,
			CreatedAt:  info.CreatedAt,
		}

		if info.DealFakeId != nil {
			response.DealId = info.DealFakeId
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(response))
	}
}
