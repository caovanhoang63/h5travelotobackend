package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/tokenprovider"
	"h5travelotobackend/component/tokenprovider/jwt"
	userstorage "h5travelotobackend/module/users/storage"
	"strings"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

// extractTokenFromHeaderString returns the access token in authorization field in request header
func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//Authorization : Bearn{token}
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user_id to find from DB

func RequireAuth(appCtx appContext.AppContext) func(ctx *gin.Context) {

	tokenProvider := jwt.NewJWTProvider(appCtx.GetSecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetGormDbConnection()

		store := userstorage.NewSqlStore(db)

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		if common.IsDebug {
			if payload.Expiry != common.AccessTokenAliveTime {
				panic(tokenprovider.ErrInvalidToken)
			}
		}

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			if err == common.RecordNotFound {
				panic(common.ErrNoPermission(errors.New("user not found")))
			}
			//c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}

}
