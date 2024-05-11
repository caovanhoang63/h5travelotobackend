package ginhotelsearch

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelsearchbiz "h5travelotobackend/search/module/hotel/biz"
	hotelmodel "h5travelotobackend/search/module/hotel/model"
	hotelsearchrepo "h5travelotobackend/search/module/hotel/repo"
	hotelstorage "h5travelotobackend/search/module/hotel/storage/esstore"
	rtSearchlocalHdl "h5travelotobackend/search/module/roomtype/transport/rtlocalhandler"
	"net/http"
)

func ListHotel(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter hotelmodel.Filter
		err := c.ShouldBindQuery(&filter)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		err = c.ShouldBindQuery(&paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.FullFill()
		store := hotelstorage.NewESStore(appCtx.GetElasticSearchClient())
		rtHandler := rtSearchlocalHdl.NewListAvailableRoomTypeHandler(appCtx)
		repo := hotelsearchrepo.NewListHotelRepo(store, rtHandler)
		biz := hotelsearchbiz.NewListHotelBiz(repo)

		result, err := biz.ListHotelWithFilter(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
