package ginhotelsearch

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelsearchbiz "h5travelotobackend/search/module/hotel/biz"
	hotelstorage "h5travelotobackend/search/module/hotel/storage/esstore"
	rtsearchstorage "h5travelotobackend/search/module/roomtype/storage"
	"net/http"
)

func ListProminentHotels(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit := c.GetInt("limit")
		if limit == 0 {
			limit = 5
		}
		store := hotelstorage.NewESStore(appCtx.GetElasticSearchClient())
		rStore := rtsearchstorage.NewStore(appCtx.GetElasticSearchClient(), appCtx.GetRedisClient())
		biz := hotelsearchbiz.NewListProminentHotelBiz(store, rStore)
		hotels, err := biz.ListProminentHotel(c.Request.Context(), limit)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(hotels))
	}
}
