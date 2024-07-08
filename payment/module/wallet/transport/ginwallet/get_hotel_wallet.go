package ginwallet

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	walletbiz "h5travelotobackend/payment/module/wallet/biz"
	walletstorage "h5travelotobackend/payment/module/wallet/storage"
	"net/http"
)

func GetHotelWalletByHotelId(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := walletstorage.NewStore(appCtx.GetGormDbConnection())
		biz := walletbiz.NewGetHotelWalletBiz(store)
		wallet, err := biz.GetHotelWalletById(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}
		wallet.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(wallet))
	}
}
