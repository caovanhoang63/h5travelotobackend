package ginroomfacilities

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	roomfacilitiesbiz "h5travelotobackend/module/roomfacilities/biz"
	roomfacilityrepo "h5travelotobackend/module/roomfacilities/repo"
	roomfacilitysqlstore "h5travelotobackend/module/roomfacilities/storage/sqlstore"
	"net/http"
)

func ListAllRoomFacilities(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := roomfacilitysqlstore.NewSqlStore(appCtx.GetGormDbConnection())
		repo := roomfacilityrepo.NewListFacilityTypesRepo(store, store)
		biz := roomfacilitiesbiz.NewListRoomFacilities(repo)
		data, err := biz.ListAllRoomFacilities(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
