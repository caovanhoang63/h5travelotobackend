package ginworker

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	workerbiz "h5travelotobackend/module/worker/biz"
	workermodel "h5travelotobackend/module/worker/model"
	workersqlstorage "h5travelotobackend/module/worker/storage/sqlstorage"
	"net/http"
)

func ListHotelWorker(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		hotelUid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.FullFill()

		filter := workermodel.Filter{HotelId: int(hotelUid.GetLocalID())}

		store := workersqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := workerbiz.NewListWorkerBiz(store)

		data, err := biz.GetHotelWorkers(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))

	}
}
