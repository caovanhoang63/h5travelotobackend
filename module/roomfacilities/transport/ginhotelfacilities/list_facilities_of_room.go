package ginroomfacilities

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	roomfacilitiesbiz "h5travelotobackend/module/roomfacilities/biz"
	roomfacilitysqlstore "h5travelotobackend/module/roomfacilities/storage/sqlstore"
	"net/http"
)

func GetFacilitiesOfRoomType(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("room-type-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := roomfacilitysqlstore.NewSqlStore(appCtx.GetGormDbConnection())
		biz := roomfacilitiesbiz.NewListFaciOfRoomBiz(store)

		data, err := biz.ListFacilitiesOfRoom(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
