package gindistrict

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	districtbiz "h5travelotobackend/module/districts/biz"
	dictrictsqlstorage "h5travelotobackend/module/districts/storage"
	"net/http"
	"strconv"
)

func ListDistrictsByProvinceCode(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		provinceCode, err := strconv.Atoi(c.Param("province-code"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := dictrictsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := districtbiz.NewListDistrictBiz(store)
		districts, err := biz.ListDistrictByProvinceCode(c.Request.Context(), provinceCode)
		if err != nil {
			panic(err)
		}

		response := common.SimpleSuccessResponse(districts)
		c.JSON(http.StatusOK, response)
		c.Set("response", response)
	}
}
