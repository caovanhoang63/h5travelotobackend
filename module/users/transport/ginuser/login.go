package ginuser

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/hasher"
	"h5travelotobackend/component/tokenprovider/jwt"
	userbiz "h5travelotobackend/module/users/business"
	usermodel "h5travelotobackend/module/users/model"
	userstorage "h5travelotobackend/module/users/storage"
	workersqlstorage "h5travelotobackend/module/worker/storage/sqlstorage"
	"net/http"
)

func UserLogin(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userLogin usermodel.UserLogin
		err := c.ShouldBind(&userLogin)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetGormDbConnection()
		tokenProvider := jwt.NewJWTProvider(appCtx.GetSecretKey()) //appctx.SecretKey()
		store := userstorage.NewSqlStore(db)
		sha256Hasher := hasher.NewSha256Hash()
		workerStorage := workersqlstorage.NewSqlStore(db)
		biz := userbiz.NewLoginBiz(appCtx, store, tokenProvider, sha256Hasher, common.AccessTokenAliveTime, common.RefreshTokenAliveTime, workerStorage)

		account, err := biz.Login(c.Request.Context(), &userLogin)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
