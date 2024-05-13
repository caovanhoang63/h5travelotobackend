package ginward

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	wardbiz "h5travelotobackend/module/ward/biz"
	wardstorage "h5travelotobackend/module/ward/districtstorage"
	"net/http"
	"strconv"
)

func ListWardsByDistrictCode(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		districtCode, err := strconv.Atoi(c.Param("district-code"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := wardstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := wardbiz.NewListWardBiz(store)
		wards, err := biz.ListWardsByDistrictCode(c.Request.Context(), districtCode)
		if err != nil {
			panic(err)
		}

		response := common.SimpleSuccessResponse(wards)
		c.JSON(http.StatusOK, response)
		c.Set("response", response)
	}
}
