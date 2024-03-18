package ginprovinces

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/module/provinces/biz"
	provincestorage "h5travelotobackend/module/provinces/storage"
	"net/http"
)

func ListAllProvinces(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := provincestorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := biz.NewListAllProvincesBiz(store)
		provinces, err := biz.ListAllProvinces(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(provinces))

	}
}
