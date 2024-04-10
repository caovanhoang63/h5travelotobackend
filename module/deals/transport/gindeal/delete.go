package gindeal

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	dealbiz "h5travelotobackend/module/deals/biz"
	dealsqlstorage "h5travelotobackend/module/deals/storage"
	"net/http"
)

func DeleteDealById(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("deal-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := dealsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := dealbiz.NewDeleteDealBiz(store)
		if err := biz.DeleteDeal(context.Background(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
