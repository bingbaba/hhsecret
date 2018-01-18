package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bingbaba/hhsecret"
)

const (
	FMT_HRSIGN_SIGNLIST_URL = "https://m.bingbaba.com/api/user/%s/sign"
)

var (
	HttpClient *http.Client
)

func init() {
	HttpClient = http.DefaultClient
}

type ListSignResp struct {
	Code int                     `json:"code"`
	Msg  string                  `json:"msg"`
	Data *hhsecret.ListSingnData `json:"data"`
}

func GetListSign(uid string) (*hhsecret.ListSingnData, error) {
	real_url := fmt.Sprintf(FMT_HRSIGN_SIGNLIST_URL, uid)
	resp, err := HttpClient.Get(real_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	lsResp := new(ListSignResp)
	err = json.Unmarshal(data, lsResp)
	if err != nil {
		return nil, err
	}

	if lsResp.Code != 1000 {
		return nil, errors.New(lsResp.Msg)
	}

	return lsResp.Data, nil
}
