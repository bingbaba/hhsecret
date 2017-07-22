package hhsecret

import (
	"fmt"
	"github.com/garyburd/go-oauth/oauth"
	"net/http"
	"net/url"
	"strings"
)

var (
	HOST_HDXT = "http://hdxt.haier.net/"
)

type Client struct {
	LoginInfo   *LoginInfo
	LoginData   *LoginData
	oauthClient oauth.Client
}

func NewClient(username, password string, consumer_key, consumer_secret string) *Client {
	return &Client{
		LoginInfo: NewLoginInfo(username, password),
		oauthClient: oauth.Client{
			Credentials: oauth.Credentials{
				Token:  consumer_key,
				Secret: consumer_secret,
			},
			SignatureMethod: oauth.HMACSHA1,
		},
	}
}

func (client *Client) Login() error {
	data, err := client.LoginInfo.Do()
	if err != nil {
		return err
	}
	client.LoginData = data
	return nil
}

func (client *Client) newHttpReq(method string, path string, form url.Values) (*http.Request, error) {
	req, err := http.NewRequest(method, HOST_HDXT, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.URL.Opaque = path

	tokenCre := &oauth.Credentials{
		Token:  client.LoginData.OauthToken,
		Secret: client.LoginData.OauthTokenSecret,
	}
	if err := client.oauthClient.SetAuthorizationHeader(req.Header, tokenCre, method, req.URL, form); err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("User-Agent",
		fmt.Sprintf("deviceid:%s;os:Andorid;osversion:7.0;manufacturer:HUAWEIMHA;model:%s;network:WIFI;spn:;appversion:4.6.8(v312);screen:1920x1200;",
			client.LoginInfo.devid,
			client.LoginInfo.devtype,
		),
	)

	return req, nil
}
