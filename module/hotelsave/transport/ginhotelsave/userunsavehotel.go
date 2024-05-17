package ginhotelsave

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	htsavebiz "h5travelotobackend/module/hotelsave/biz"
	htsavemodel "h5travelotobackend/module/hotelsave/model"
	htsavestore "h5travelotobackend/module/hotelsave/store"
	"net/http"
)

func UserUnSaveHotel(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		uid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := htsavestore.NewStore(appCtx.GetGormDbConnection())
		biz := htsavebiz.NewUnsaveHotelBiz(store)

		del := &htsavemodel.HotelSaveDelete{
			UserId:  requester.GetUserId(),
			HotelId: int(uid.GetLocalID()),
		}
		if err = biz.UnsaveHotel(c.Request.Context(), del); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
