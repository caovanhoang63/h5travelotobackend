package ginpayin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	payinbiz "h5travelotobackend/payment/module/payin/biz"
	payinmodel "h5travelotobackend/payment/module/payin/model"
	payinstore "h5travelotobackend/payment/module/payin/store"
	"net/http"
)

func CancelPayIn(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		txnId := c.Param("txn_id")
		if txnId == "" {
			panic(common.ErrInvalidRequest(errors.New("txn_id not found")))
		}
		update := payinmodel.PaymentBookingUpdateStatus{PaymentStatus: common.PaymentStatusFailed}

		store := payinstore.NewStore(appCtx.GetGormDbConnection())
		biz := payinbiz.NewUpdatePBStatusBiz(store)

		if err := biz.UpdateStatus(c.Request.Context(), txnId, &update); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
