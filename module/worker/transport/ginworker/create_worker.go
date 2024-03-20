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

func CreateWorker(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var worker workermodel.WorkerCreate
		hotelUid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&worker); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		worker.HotelId = int(hotelUid.GetLocalID())
		worker.UnMask()

		store := workersqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := workerbiz.NewCreateWorkerBiz(store, store)

		if err := biz.CreateWorker(c.Request.Context(), &worker); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
