package ginhotelfacilities

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelfacilitiesbiz "h5travelotobackend/module/hotelfacilities/biz"
	hotelfacilitiesstorage "h5travelotobackend/module/hotelfacilities/storage"
	"net/http"
)

func ListAllHotelFacilities(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := hotelfacilitiesstorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := hotelfacilitiesbiz.NewListHotelFacilities(store)
		data, err := biz.ListAllHotelFacilities(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}
