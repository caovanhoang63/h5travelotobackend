package ginuser

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	userbiz "h5travelotobackend/module/users/business"
	usermodel "h5travelotobackend/module/users/model"
	userstorage "h5travelotobackend/module/users/storage"
)

func ConfirmResetPassword(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")

		if email == "" {
			panic(common.ErrInvalidRequest(usermodel.ErrInvalidEmail))
		}

		store := userstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := userbiz.NewConfirmResetPassword(store, appCtx.GetCacher(), appCtx.GetSendMailEngine())
		err := biz.ConfirmResetPassword(c.Request.Context(), email)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))

	}
}
