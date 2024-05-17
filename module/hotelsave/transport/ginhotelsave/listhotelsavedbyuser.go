package ginhotelsave

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	htsavebiz "h5travelotobackend/module/hotelsave/biz"
	htsavestore "h5travelotobackend/module/hotelsave/store"
	"net/http"
)

func ListHotelSavedByUser(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.FullFill()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := htsavestore.NewStore(appCtx.GetGormDbConnection())
		biz := htsavebiz.NewListHotelSavedBiz(store)

		var data []common.Hotel

		data, err := biz.ListHotelsSaveByUser(c.Request.Context(), nil, &paging, requester)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}
