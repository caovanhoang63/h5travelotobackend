package ginpayin

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	paymentmodel "h5travelotobackend/payment/module/payin/model"
	"net/http"
)

func PayIn(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var info paymentmodel.PaymentInfo

		if err := c.ShouldBind(&info); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := info.UnMask(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		vnPay := appCtx.GetVnPay()

		url := vnPay.NewPayInUrl(100000, info.BookingFakeId.String(), c.ClientIP(), "123123124")

		response := paymentmodel.PaymentInfoResponse{
			PaymentUrl: url,
			Amount:     1000000,
			Currency:   common.VND,
			BookingId:  info.BookingFakeId.String(),
			Method:     common.PaymentMethodVnPay,
		}

		if info.DealFakeId != nil {
			response.DealId = info.DealFakeId.String()
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(response))
	}
}
