package gindistrict

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	districtbiz "h5travelotobackend/module/districts/biz"
	"h5travelotobackend/module/districts/districtstorage"
	"net/http"
	"strconv"
)

func ListDistrictsByProvinceCode(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		provinceCode, err := strconv.Atoi(c.Param("province-code"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := districtstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := districtbiz.NewListDistrictBiz(store)
		districts, err := biz.ListDistrictByProvinceCode(c.Request.Context(), provinceCode)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(districts))
	}
}
