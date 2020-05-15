package web

import (
	"github.com/gin-gonic/gin"
)

func GetApp() (app *gin.Engine) {
	app = gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	// 用户API
	userApi := app.Group("/api/user/:username", UserWhiteList)
	{
		// query
		userApi.POST("/login", UserLoginHander)
		userApi.GET("/login", UserLoginCheckHander)
		userApi.POST("/sign", UserSignHander)
		userApi.GET("/sign", UserListSignHander)
		userApi.GET("/sign/month/:year/:month", UserMonthSignHandler)

		// notice
		userApi.GET("/notice", NoticeHander)
	}

	// static web
	app.Static("/html", DefaultCfg.StaticWebPath)

	return
}
