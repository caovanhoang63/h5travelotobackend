package gindeal

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	dealbiz "h5travelotobackend/module/deals/biz"
	dealsqlstorage "h5travelotobackend/module/deals/storage"
	"net/http"
)

func GetDealById(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("deal-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := dealsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := dealbiz.NewFindDealBiz(store)
		deal, err := biz.FindDealById(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		deal.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(deal))

	}
}
