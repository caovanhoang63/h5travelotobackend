package gindeal

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	dealbiz "h5travelotobackend/module/deals/biz"
	dealmodel "h5travelotobackend/module/deals/model"
	dealsqlstorage "h5travelotobackend/module/deals/storage"
	"net/http"
)

func UpdateDeal(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("deal-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data dealmodel.DealUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := dealsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := dealbiz.NewUpdateDealBiz(store)
		if err := biz.UpdateDeal(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
