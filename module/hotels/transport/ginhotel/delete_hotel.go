package ginhotel

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelbiz "h5travelotobackend/module/hotels/business"
	hotelrepo "h5travelotobackend/module/hotels/repo"
	hotelmongostorage "h5travelotobackend/module/hotels/storage/mongostorage"
	hotelmysqlstorage "h5travelotobackend/module/hotels/storage/mysqlstorage"
	"net/http"
)

func DeleteHotel(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		uid, err := common.FromBase58(context.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := context.MustGet(common.CurrentUser).(common.Requester)

		mongoS := hotelmongostorage.NewMongoStore(appCtx.GetMongoConnection())
		sqlS := hotelmysqlstorage.NewSqlStore(appCtx.GetGormDbConnection())

		deleteRepo := hotelrepo.NewDeleteHotelRepo(mongoS, sqlS)
		findRepo := hotelrepo.NewFindHotelRepo(sqlS, mongoS)
		biz := hotelbiz.NewDeleteHotelBiz(deleteRepo, findRepo, requester)

		if err := biz.DeleteHotel(context.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
