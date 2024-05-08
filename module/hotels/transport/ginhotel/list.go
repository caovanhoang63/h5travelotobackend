package ginhotel

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelbiz "h5travelotobackend/module/hotels/biz"
	hotelmodel "h5travelotobackend/module/hotels/model"
	hotelstorage "h5travelotobackend/module/hotels/storage"
	reviewstorage "h5travelotobackend/module/review/storage/redis"
)

func ListHotel(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter hotelmodel.Filter
		var paging common.Paging
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.FullFill()

		store := hotelstorage.NewSqlStore(appCtx.GetGormDbConnection())
		rvStore := reviewstorage.NewRedisStore(appCtx.GetRedisClient())
		biz := hotelbiz.NewListHotelBiz(store, rvStore)
		data, err := biz.ListHotel(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}
		c.JSON(200, common.NewSuccessResponse(data, paging, filter))
	}
}
