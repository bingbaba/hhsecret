package web

import (
	"github.com/gin-gonic/gin"
)

func GetApp() (app *gin.Engine) {
	app = gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	// query
	app.POST("/api/user/:username/login", UserWhiteList, UserLoginHander)
	app.GET("/api/user/:username/login", UserWhiteList, UserLoginCheckHander)
	app.POST("/api/user/:username/sign", UserWhiteList, UserSignHander)
	app.GET("/api/user/:username/sign", UserWhiteList, UserListSignHander)
	app.GET("/api/user/:username/sign/month/:year/:month", UserWhiteList, UserMonthSignHandler)

	// notice
	app.GET("/api/user/:username/notice", UserWhiteList, NoticeHander)

	// static web
	app.Static("/html", DefaultCfg.StaticWebPath)

	return
}
