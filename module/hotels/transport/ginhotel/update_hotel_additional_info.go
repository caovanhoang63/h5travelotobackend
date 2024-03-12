package ginhotel

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelbiz "h5travelotobackend/module/hotels/business"
	hotelmodel "h5travelotobackend/module/hotels/model"
	hotelrepo "h5travelotobackend/module/hotels/repo"
	hotelmongostorage "h5travelotobackend/module/hotels/storage/mongostorage"
	hotelmysqlstorage "h5travelotobackend/module/hotels/storage/mysqlstorage"
	"net/http"
)

func UpdateHotelAdditionalInfo(appCtx appContext.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data hotelmodel.HotelAdditionalInfo

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		sqlStore := hotelmysqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		mongoStore := hotelmongostorage.NewMongoStore(appCtx.GetMongoConnection())
		findRepo := hotelrepo.NewFindHotelRepo(sqlStore, mongoStore)
		biz := hotelbiz.NewHotelAdditionalDataUpdateBiz(mongoStore, findRepo, requester)

		if err := biz.UpdateAdditionalInfo(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
