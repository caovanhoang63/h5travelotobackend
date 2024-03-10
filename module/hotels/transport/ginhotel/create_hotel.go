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

func CreateHotel(appCtx appContext.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data hotelmodel.HotelCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data.OwnerID = requester.GetUserId()

		sqlStore := hotelmysqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		mongoStore := hotelmongostorage.NewMongoStore(appCtx.GetMongoConnection())
		repo := hotelrepo.NewCreateHotelRepo(sqlStore, mongoStore)
		biz := hotelbiz.NewCreateHotelBiz(repo)

		if err := biz.CreateHotel(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))

	}
}
