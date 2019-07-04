package web

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
	"time"

	"github.com/bingbaba/hhsecret"
)

var (
	ERROR_USER_NOTALLOW  = errors.New("user not allow")
	ERROR_USER_NOTLOGIN  = errors.New("user not login")
	ERROR_REQUIRE_PASSWD = errors.New("require password to login")
)

func UserWhiteList(ctx *gin.Context) {
	username := ctx.Param("username")
	for _, un_tmp := range DefaultCfg.UserWhiteList {
		if username == un_tmp {
			ctx.Next()
			return
		}
	}

	ctx.JSON(401, NewResponseWithErr(ERROR_USER_NOTALLOW, nil))
}

func UserLoginCheckHander(ctx *gin.Context) {
	var result = make(map[string]string)
	username := ctx.Param("username")
	client, found := GetClientByUser(username)
	if !found {
		ctx.JSON(200, NewResponseWithErr(ERROR_USER_NOTLOGIN, result))
	} else {
		ctx.JSON(200, NewResponse(client.LoginData))
	}
}

func UserLoginHander(ctx *gin.Context) {
	var err error
	var result interface{}
	username := ctx.Param("username")
	defer func() {
		if err != nil {
			logger.Errorf("%s login failed:", username, err)
		}
		ctx.JSON(200, NewResponseWithErr(err, result))
	}()

	userInfo := make(map[string]string)
	err = ctx.BindJSON(&userInfo)
	if err != nil {
		body, err2 := ioutil.ReadAll(ctx.Request.Body)
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

func UserSignHander(ctx *gin.Context) {
	var err error
	var result interface{}
	username := ctx.Param("username")
	defer func() {
		if err != nil {
			logger.Errorf("%s sign failed: %v", username, err)
		}
		ctx.JSON(200, NewResponseWithErr(err, result))
	}()

	client, found := GetClientByUser(username)
	if !found {
		err = ERROR_USER_NOTLOGIN
		return
	}

	result, err = client.Sign()
	return
}
func UserListSignHander(ctx *gin.Context) {
	var err error
	var result interface{}
	username := ctx.Param("username")
	defer func() {
		ctx.JSON(200, NewResponseWithErr(err, result))
	}()

	client, found := GetClientByUser(username)
	if !found {
		err = ERROR_USER_NOTLOGIN
		return
	}

	result, err = client.ListSignPost()
	return
}

func UserMonthSignHandler(ctx *gin.Context) {
	username := ctx.Param("username")
	year := ctx.Param("year")
	month := ctx.Param("month")

	var ms *hhsecret.MonthSign
	var err error
	defer func() {
		ctx.JSON(200, NewResponseWithErr(err, ms))
	}()
	ms, err = hhsecret.GetMonthSign(username, year, month)
	return
}

func NoticeHander(ctx *gin.Context) {
	username := ctx.Param("username")

	var notice = false
	var afternoon = false
	if time.Now().Hour() >= 12 {
		afternoon = true

		// not send notice before 17:00
		if time.Now().Hour() < 17 {
			notice = false
			ctx.JSON(200, NewResponse(notice))
			return
		}
	}

	client, found := GetClientByUser(username)
	if !found {
		ctx.JSON(200, NewResponseWithErr(errors.New("user not login"), false))
		return
	}

	lsd, err := client.ListSignPost()
	if err != nil {
		ctx.JSON(200, NewResponseWithErr(err, nil))
	} else {
		if len(lsd.Signs) == 0 {
			notice = true
		} else {
			if afternoon {
				mtime := lsd.Signs[0].GetMinuteSecode()
				if strings.Compare(mtime, "17:30") < 0 {
					notice = true
				}
			}
		}

		ctx.JSON(200, NewResponse(notice))
	}
}
