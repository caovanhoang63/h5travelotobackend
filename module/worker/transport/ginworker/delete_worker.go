package ginworker

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	workerbiz "h5travelotobackend/module/worker/biz"
	workersqlstorage "h5travelotobackend/module/worker/storage/sqlstorage"
	"net/http"
)

func DeleteWorker(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		hotelUid, err := common.FromBase58(c.Param("hotel-id"))
		userUid, err := common.FromBase58(c.Param("user-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := workersqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := workerbiz.NewDeleteWorkerBiz(store, store)
		if err := biz.DeleteWorker(c.Request.Context(), int(hotelUid.GetLocalID()), int(userUid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
