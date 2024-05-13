package ginrtsearch

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/module/bookings/transport/bklocalhandler"
	rtsearchbiz "h5travelotobackend/search/module/roomtype/biz"
	rtsearchmodel "h5travelotobackend/search/module/roomtype/model"
	rtsearchrepo "h5travelotobackend/search/module/roomtype/repo"
	rtsearchstorage "h5travelotobackend/search/module/roomtype/storage"
	"h5travelotobackend/search/utils"
	"log"
	"net/http"
)

func ListAvailableRoomType(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter rtsearchmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter.UnMask(false)

		rtStore := rtsearchstorage.NewStore(appCtx.GetElasticSearchClient(), appCtx.GetRedisClient())
		bkHandler := bklocalhandler.NewCountBookedRoomLocalHandler(appCtx)
		rtHandlerRepo := rtsearchrepo.NewListRoomTypeRepo(rtStore, bkHandler)
		biz := rtsearchbiz.NewListAvailableRtBiz(rtHandlerRepo)

		if filter.QueryTime != 0 {
			key := utils.GenCacheKeyForQuery(filter.QueryTime, filter.HotelId)
			log.Println("key: ", key)
			rts, err := rtStore.GetAvailableRoomTypeInCache(c.Request.Context(), key)
			if err != nil {
				log.Println("error cache ", err)
			}
			if rts != nil && err == nil {
				log.Println("cache ")
				c.Data(http.StatusOK, "application/json", rts)
				return
			}
		}

		filter.SetDefault()
		rts, err := biz.ListAvailableRt(c.Request.Context(), &filter)
		if err != nil {
			panic(err)
		}
		for i := range rts {
			rts[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(rts, nil, filter))
	}
}
