package ginwallet

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	walletbiz "h5travelotobackend/payment/module/wallet/biz"
	walletmodel "h5travelotobackend/payment/module/wallet/model"
	walletstorage "h5travelotobackend/payment/module/wallet/storage"
	"net/http"
)

func Withdrawal(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data walletmodel.HotelWalletWithdrawal
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		data.UnMask()

		store := walletstorage.NewStore(appCtx.GetGormDbConnection())
		biz := walletbiz.NewWithdrawalBiz(store)

		err = biz.Withdrawal(c.Request.Context(), int(uid.GetLocalID()), data.Amount)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
