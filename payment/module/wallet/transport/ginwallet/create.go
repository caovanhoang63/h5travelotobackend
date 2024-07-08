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

func CreateWallet(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var wallet walletmodel.HotelWalletCreate
		if err := c.ShouldBindJSON(&wallet); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		wallet.UnMask()
		store := walletstorage.NewStore(appCtx.GetGormDbConnection())
		biz := walletbiz.NewCreateWalletBiz(store)
		err := biz.CreateWallet(c.Request.Context(), &wallet)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
