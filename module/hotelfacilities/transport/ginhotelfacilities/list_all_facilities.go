package ginhotelfacilities

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelfacilitiesbiz "h5travelotobackend/module/hotelfacilities/biz"
	hotelfacilityrepo "h5travelotobackend/module/hotelfacilities/repo"
	hotelfacilitysqlstore "h5travelotobackend/module/hotelfacilities/storage/sqlstore"
	"net/http"
)

func ListAllHotelFacilities(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := hotelfacilitysqlstore.NewSqlStore(appCtx.GetGormDbConnection())
		repo := hotelfacilityrepo.NewListFacilityTypesRepo(store, store)
		biz := hotelfacilitiesbiz.NewListHotelFacilities(repo)
		data, err := biz.ListAllHotelFacilities(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
