package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

const (
	FMT_HRSIGN_NOTICECHECK_URL = "https://m.bingbaba.com/api/user/%s/notice"
)

type SignNoticeResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Notice bool   `json:"data"`
}

func IfNotice(uid string) (bool, error) {
	real_url := fmt.Sprintf(FMT_HRSIGN_NOTICECHECK_URL, uid)
	resp, err := HttpClient.Get(real_url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	snResp := new(SignNoticeResp)
	err = json.Unmarshal(data, snResp)
	if err != nil {
		return false, fmt.Errorf("parse %s failed:%v", data, err)
	}

	if snResp.Code != 1000 {
		return false, errors.New(snResp.Msg)
	}

	return snResp.Notice, nil
}
