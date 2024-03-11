package main

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/middleware"
	"h5travelotobackend/module/hotels/transport/ginhotel"
	"h5travelotobackend/module/users/transport/ginuser"
)

func SetUpRoute(appCtx appContext.AppContext, v1 *gin.RouterGroup) {
	v1.POST("/register", ginuser.RegisterUser(appCtx))
	v1.POST("/authenticate", ginuser.UserLogin(appCtx))

	users := v1.Group("/users", middleware.RequireAuth(appCtx))

	users.GET("/profile", ginuser.GetProfile(appCtx))
	users.PATCH("/profile", ginuser.Update(appCtx))
	users.PATCH("/change-password", ginuser.ChangePassword(appCtx))

	hotels := v1.Group("hotels", middleware.RequireAuth(appCtx))
	hotels.POST("/", ginhotel.CreateHotel(appCtx))
	hotels.DELETE("/:id", ginhotel.DeleteHotel(appCtx))

}
