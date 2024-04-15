package ginuser

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	userbiz "h5travelotobackend/module/users/business"
	userstorage "h5travelotobackend/module/users/storage"
	"net/http"
)

func CheckExistedEmail(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var email string
		email = c.Query("email")

		store := userstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := userbiz.NewCheckExistedEmailBiz(store)

		if err := biz.CheckExistedEmail(c.Request.Context(), email); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(false))
	}
}
