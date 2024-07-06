package ginhotel

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelbiz "h5travelotobackend/module/hotels/biz"
	hotelstorage "h5travelotobackend/module/hotels/storage"
	"net/http"
)

func GetHotelByUser(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		worker := c.MustGet(common.CurrentWorker).(common.Worker)
		store := hotelstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := hotelbiz.NewFindHotelBiz(store, appCtx.GetPubSub())
		data, err := biz.FindWithConditionHotel(c.Request.Context(), map[string]interface{}{"id": worker.GetHotelId()})
		if err != nil {
			panic(err)
		}
		data.Mask(false)
		response := common.SimpleSuccessResponse(data)
		c.JSON(http.StatusOK, response)
	}
}
