package client

import (
	"github.com/bingbaba/hhsecret"
)

func Login(username, password string) (*hhsecret.LoginData, error) {
	login := hhsecret.NewLoginInfo(username, password)
	return login.Do()
}
