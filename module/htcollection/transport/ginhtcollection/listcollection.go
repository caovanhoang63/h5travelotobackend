package ginhtcollection

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	htcollectionbiz "h5travelotobackend/module/htcollection/biz"
	htcollection "h5travelotobackend/module/htcollection/model"
	htcollectionstore "h5travelotobackend/module/htcollection/store"
	"net/http"
)

func ListUserCollections(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.FullFill()
		var filter htcollection.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := filter.UnMask(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := htcollectionstore.NewStore(appCtx.GetGormDbConnection())
		biz := htcollectionbiz.NewListCollectionBiz(store)
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		collections, err := biz.ListCollectionsWithCondition(c.Request.Context(), &filter, &paging, requester)
		if err != nil {
			panic(err)
		}

		for i := range collections {
			collections[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(collections, paging, filter))
	}
}
