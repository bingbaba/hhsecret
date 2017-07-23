package hhsecret

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var (
	random *rand.Rand
)

func init() {
	random = rand.New(rand.NewSource(time.Now().Unix()))
}

type ListSignResp struct {
	Success   bool           `json:"success"`
	Error     string         `json:"errorMessage"`
	ErrorCode int            `json:"errorCode"`
	Data      *ListSingnData `json:"data"`
}

type ListSingnData struct {
	Signs     []*Sign `json:"signs"`
	TimeStamp int64   `json:"time"`
	Count     int     `json:"count"`
}

type Sign map[string]interface{}

func (client *Client) Sign() (*ListSingnData, error) {
	url_param := url.Values{
		"time":   {time.Now().Format("2006-01-02 15:04:05")},
		"source": {client.oauthClient.Credentials.Token}, //consumer_key
	}
	path := fmt.Sprintf("/snsapi/%s/attendance/sign.json?"+url_param.Encode(), client.LoginData.OrgInfoId)

	lat := fmt.Sprintf("%0.6f", 36.130+random.Float64()/1000)
	lng := fmt.Sprintf("%0.6f", 120.416+random.Float64()/1000)
	form := url.Values{
		"newFlag":   {"newFlag"},
		"account":   {client.LoginInfo.userName},
		"mid":       {"102"},
		"latitude":  {lat},
		"longitude": {lng},
		"deviceId":  {client.LoginInfo.devid},
	}
	req, err := client.newHttpReq("POST", path, form)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	lsr := &ListSignResp{}
	err = json.Unmarshal(data, lsr)
	if err != nil {
		return nil, err
	}

	if !lsr.Success {
		return nil, errors.New(lsr.Error)
	}

	return lsr.Data, nil
}

func (client *Client) ListSignPost() (*ListSingnData, error) {
	form := url.Values{
		"midpost":     {"102"},
		"accountpost": {client.LoginInfo.userName},
	}
	path := fmt.Sprintf("/snsapi/%s/attendance/list_signpost.json", client.LoginData.OrgInfoId)
	req, err := client.newHttpReq("POST", path, form)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	lsr := &ListSignResp{}
	err = json.Unmarshal(data, lsr)
	if err != nil {
		return nil, err
	}

	if !lsr.Success {
		return nil, errors.New(lsr.Error)
	}

	return lsr.Data, nil
}
