package ginhotel

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelbiz "h5travelotobackend/module/hotels/business"
	hotelmodel "h5travelotobackend/module/hotels/model"
	hotelrepo "h5travelotobackend/module/hotels/repo"
	hotelmysqlstorage "h5travelotobackend/module/hotels/storage/mysqlstorage"
	"net/http"
)

func ListHotel(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter hotelmodel.Filter
		if err := c.ShouldBindQuery(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data []hotelmodel.Hotel
		var paging common.Paging

		err := c.ShouldBind(&paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.FullFill()
		err = c.ShouldBind(&filter)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := hotelmysqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		repo := hotelrepo.NewListRestaurantRepo(store)
		biz := hotelbiz.NewListRestaurantBiz(repo)

		data, err = biz.List(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))

	}
}
