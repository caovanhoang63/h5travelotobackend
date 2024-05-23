package ginhotelsearch

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/module/hotels/storage/hotelrdbstore"
	hotelsearchbiz "h5travelotobackend/search/module/hotel/biz"
	hotelsearchrepo "h5travelotobackend/search/module/hotel/repo"
	hotelstorage "h5travelotobackend/search/module/hotel/storage/esstore"
	hotelsearchrdbstore "h5travelotobackend/search/module/hotel/storage/rdb"
	rtsearchstorage "h5travelotobackend/search/module/roomtype/storage"
	"net/http"
)

func ListRecentlyViewed(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		store := hotelstorage.NewESStore(appCtx.GetElasticSearchClient())
		rStore := rtsearchstorage.NewStore(appCtx.GetElasticSearchClient(), appCtx.GetRedisClient())
		rdbStore := hotelrdbstore.NewStore(appCtx.GetRedisClient())
		rtStore := hotelsearchrdbstore.NewStore(appCtx.GetRedisClient())
		repo := hotelsearchrepo.NewHotelsByIdsRepo(store, rStore, rtStore)
		biz := hotelsearchbiz.NewListRecentlyViewedBiz(rdbStore, repo)
		result, err := biz.ListRecentlyViewed(c.Request.Context(), requester.GetUserId())
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
