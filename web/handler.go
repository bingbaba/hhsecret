package web

import (
	"errors"
	"github.com/bingbaba/hhsecret"
	"github.com/kataras/iris/context"
	"io/ioutil"
)

var (
	ERROR_USER_NOTALLOW  = errors.New("user not allow")
	ERROR_USER_NOTLOGIN  = errors.New("user not login")
	ERROR_REQUIRE_PASSWD = errors.New("require password to login")
)

func UserWhiteList(ctx context.Context) {
	username := ctx.Params().Get("username")
	for _, un_tmp := range DefaultCfg.UserWhiteList {
		if username == un_tmp {
			ctx.Next()
			return
		}
	}

	ctx.StatusCode(401)
	ctx.JSON(NewResponseWithErr(ERROR_USER_NOTALLOW, nil))
}

func UserLoginHander(ctx context.Context) {
	var err error
	var result interface{}
	username := ctx.Params().Get("username")
	defer func() {
		ctx.JSON(NewResponseWithErr(err, result))
	}()

	userInfo := make(map[string]string)
	err = ctx.ReadJSON(&userInfo)
	if err != nil {
		body, err2 := ioutil.ReadAll(ctx.Request().Body)
		if err2 != nil {
			err = err2
			return
		}
		if len(body) != 0 {
			return
		}
	}

	var found bool
	var client *hhsecret.Client
	password, found := userInfo["password"]
	if !found {
		client, found = GetClientByUser(username)
		if !found {
			err = ERROR_REQUIRE_PASSWD
			return
		}
	} else {
		client = hhsecret.NewClient(username, password, DefaultCfg.ConsumerKey, DefaultCfg.ConsumerSecret)
		err = client.Login()
		if err != nil {
			return
		}
	}
	result = client.LoginData
	SaveClient(username, client)

	return
}

func UserSignHander(ctx context.Context) {
	var err error
	var result interface{}
	username := ctx.Params().Get("username")
	defer func() {
		ctx.JSON(NewResponseWithErr(err, result))
	}()

	client, found := GetClientByUser(username)
	if !found {
		err = ERROR_USER_NOTLOGIN
		return
	}

	result, err = client.Sign()
	return
}
func UserListSignHander(ctx context.Context) {
	var err error
	var result interface{}
	username := ctx.Params().Get("username")
	defer func() {
		ctx.JSON(NewResponseWithErr(err, result))
	}()

	client, found := GetClientByUser(username)
	if !found {
		err = ERROR_USER_NOTLOGIN
		return
	}

	result, err = client.ListSignPost()
	return
}
