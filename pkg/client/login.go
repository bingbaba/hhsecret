package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bingbaba/hhsecret"
)

const (
	FMT_HRSIGN_LOGIN_URL = "https://m.bingbaba.com/api/user/%s/login"
)

type LoginResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Login(username, password string) error {
	req_url := fmt.Sprintf(FMT_HRSIGN_LOGIN_URL, username)

	req_map := map[string]string{"username": username, "password": password}
	body_bytes, err := json.Marshal(req_map)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, req_url, bytes.NewReader(body_bytes))
	if err != nil {
		return err
	}

	resp, err := HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	lResp := new(LoginResp)
	err = json.Unmarshal(data, lResp)
	if err != nil {
		return fmt.Errorf("parse %s failed:%v", data, err)
	}

	if lResp.Code != 1000 {
		return errors.New(lResp.Msg)
	}

	return nil
}

type LoginInfoResp struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data *hhsecret.LoginData `json:"data"`
}

func GetLoginInfo(username string) (*hhsecret.LoginData, error) {
	req_url := fmt.Sprintf(FMT_HRSIGN_LOGIN_URL, username)
	resp, err := HttpClient.Get(req_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	lResp := new(LoginInfoResp)
	err = json.Unmarshal(data, lResp)
	if err != nil {
		return nil, fmt.Errorf("parse %s failed:%v", data, err)
	}

	if lResp.Code != 1000 {
		return nil, errors.New(lResp.Msg)
	}

	return lResp.Data, nil
}
