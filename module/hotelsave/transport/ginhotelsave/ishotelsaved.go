package ginhotelsave

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	htsavebiz "h5travelotobackend/module/hotelsave/biz"
	htsavestore "h5travelotobackend/module/hotelsave/store"
	"net/http"
)

func IsHotelSaved(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		uid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := htsavestore.NewStore(appCtx.GetGormDbConnection())
		biz := htsavebiz.NewFindSavedHotelBiz(store)

		ok := biz.IsHotelSaved(c.Request.Context(), requester, int(uid.GetLocalID()))
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(ok))
	}
}
