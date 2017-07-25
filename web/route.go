package web

import (
	"github.com/kataras/iris"
)

func GetIrisApp() (app *iris.Application) {
	app = iris.New()

	// query
	app.Post("/api/user/{username}/login", UserWhiteList, UserLoginHander)
	app.Get("/api/user/{username}/login", UserWhiteList, UserLoginCheckHander)
	app.Post("/api/user/{username}/sign", UserWhiteList, UserSignHander)
	app.Get("/api/user/{username}/sign", UserWhiteList, UserListSignHander)

	// static web
	app.StaticWeb("/html", DefaultCfg.StaticWebPath)

	return
}
