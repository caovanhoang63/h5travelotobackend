package main

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/middleware"
	"h5travelotobackend/module/users/transport/ginuser"
)

func SetUpRoute(appCtx appContext.AppContext, v1 *gin.RouterGroup) {
	v1.POST("/register", ginuser.RegisterUser(appCtx))
	v1.POST("/authenticate", ginuser.UserLogin(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx), ginuser.GetProfile(appCtx))

}
