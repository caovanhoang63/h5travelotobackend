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
	"strings"
)

func GetHotelById(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		mongoS := hotelmongostorage.NewMongoStore(appCtx.GetMongoConnection())
		sqlS := hotelmysqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		repo := hotelrepo.NewFindHotelRepo(sqlS, mongoS)
		biz := hotelbiz.NewFindHotelBiz(repo)

		add := false
		if v, ok := c.GetQuery("add"); ok {
			if strings.Compare(strings.TrimSpace(v), "true") == 0 {
				add = true
			}
		}

		data, err := biz.GetHotelById(c.Request.Context(), int(uid.GetLocalID()), add)
		if err != nil {
			panic(err)
		}

		if !add {
			//data.HotelAdditionalInfo = nil
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
