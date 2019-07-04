package hhsecret

import (
	"github.com/garyburd/go-oauth/oauth"
	"net/http"
	"net/url"
	"strings"
)

var (
	HOST_HDXT = "https://i.haier.net"
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

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", client.LoginInfo.useragent)
	//req.Header.Set("User-Agent", "38882/10.1.8(1050);Android 8.0.0;Xiaomi;MIX+2;102;1080*2030;deviceId:e4d14146-18a1-3442-a740-f907c13e0c7c;deviceName:Xiaomi MIX+2;clientId:38882;os:Android 8.0.0;brand:Xiaomi;model:MIX+2;deviceKey:99001021267878,865736035357535;oem:ihaier;lang:zh-CN;")
	req.Header.Set("opentoken", client.LoginData.Token)

	return req, nil
}
