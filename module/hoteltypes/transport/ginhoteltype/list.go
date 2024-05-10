package ginhoteltype

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hoteltypebiz "h5travelotobackend/module/hoteltypes/biz"
	hoteltypestorage "h5travelotobackend/module/hoteltypes/storage"
	"net/http"
)

func ListAllHotelTypes(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		store := hoteltypestorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := hoteltypebiz.NewListHotelTypeBiz(store)

		data, err := biz.ListAllHotelTypes(c.Request.Context())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
