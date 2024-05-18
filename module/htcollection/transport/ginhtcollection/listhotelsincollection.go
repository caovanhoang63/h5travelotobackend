package ginhtcollection

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	htcollectionbiz "h5travelotobackend/module/htcollection/biz"
	htcollectionstore "h5travelotobackend/module/htcollection/store"
	"net/http"
)

func ListHotelInCollection(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("collection-id"))
		var paging common.Paging
		if err = c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.FullFill()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := htcollectionstore.NewStore(appCtx.GetGormDbConnection())
		biz := htcollectionbiz.NewListHotelsInCollectionBiz(store)

		var data []common.Hotel

		data, err = biz.ListHotelsInCollection(c.Request.Context(), nil, int(uid.GetLocalID()), &paging, requester)
		if err != nil {
			panic(err)
		}
		for i := range data {
			data[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}
