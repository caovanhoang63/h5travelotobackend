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

func ListDeal(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var filter dealmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter.UnMask()
		paging.FullFill()

		store := dealsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := dealbiz.NewListDealBiz(store)
		deals, err := biz.ListDeal(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range deals {
			deals[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(deals, paging, filter))
	}
}
