package client

import (
	"encoding/json"
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
	req_url := fmt.Sprintf(FMT_HRSIGN_LOGIN_URL, uid)

	req_map := map[string]string{"username": "", "password": ""}
	body, err := json.Marshal(req_map)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, req_url, body)
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
