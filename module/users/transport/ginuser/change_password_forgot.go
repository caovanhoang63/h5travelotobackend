package ginuser

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/hasher"
	userbiz "h5travelotobackend/module/users/business"
	usermodel "h5travelotobackend/module/users/model"
	userstorage "h5travelotobackend/module/users/storage"
)

func ChangePasswordForgot(appCtx appContext.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		email := c.Param("email")
		if email == "" {
			panic(common.ErrInvalidRequest(usermodel.ErrInvalidEmail))
		}

		var data usermodel.UserChangePassword

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSqlStore(appCtx.GetGormDbConnection())
		sha256Hasher := hasher.NewSha256Hash()
		biz := userbiz.NewChangePasswordForgot(store, appCtx.GetCacher(), sha256Hasher)

		if err := biz.ChangePassword(c.Request.Context(), email, data.Pin, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
