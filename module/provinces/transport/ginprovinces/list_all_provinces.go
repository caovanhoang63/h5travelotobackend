package ginprovinces

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/module/provinces/provincesbiz"
	provincestorage "h5travelotobackend/module/provinces/storage"
	"net/http"
)

func ListAllProvinces(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := provincestorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := provincesbiz.NewListAllProvincesBiz(store)
		provinces, err := biz.ListAllProvinces(c.Request.Context())
		if err != nil {
			panic(err)
		}
		response := common.SimpleSuccessResponse(provinces)
		c.JSON(http.StatusOK, response)
		c.Set("response", response)
	}
}
