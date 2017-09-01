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

type SignResp struct {
	Success   bool       `json:"success"`
	Error     string     `json:"errorMessage"`
	ErrorCode int        `json:"errorCode"`
	Data      *SingnData `json:"data"`
}

type SingnData struct {
	Id            string   `json:"id"`
	AttendSetId   string   `json:"attendSetId"`
	Datetime      int64    `json:"datetime"`
	Longitude     float64  `json:"longitude"`
	Latitude      float64  `json:"latitude"`
	Featurename   string   `json:"featurename"`
	Content       string   `json:"content"`
	Status        int      `json:"status"`
	MbShare       string   `json:"mbShare"`
	PhotoIds      []string `json:"photoIds"`
	ClockInType   string   `json:"clockInType"`
	ExtraRemark   string   `json:"extraRemark"`
	Systime       string   `json:"systime"`
	MessageByTime string   `json:"messageByTime"`
	Email         string   `json:"email"`
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

type Sign struct {
	DateTime  int64   `json:"datetime"`
	Location  string  `json:"featurename"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (s *Sign) GetMinuteSecode() string {
	return time.Unix(s.DateTime/1000, 0).Format("15:04")
}

func (client *Client) Sign() (*SingnData, error) {
	// url_param := url.Values{
	// 	"time":   {time.Now().Format("20060102150405")},
	// 	"source": {client.oauthClient.Credentials.Token}, //consumer_key
	// }
	// path := fmt.Sprintf("/snsapi/%s/attendance/sign.json?"+url_param.Encode(), client.LoginData.OrgInfoId)
	path := fmt.Sprintf("/snsapi/%s/attendance/sign.json", client.LoginData.OrgInfoId)

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
	// fmt.Printf("%s\n", data)

	sr := &SignResp{}
	err = json.Unmarshal(data, sr)
	if err != nil {
		err = fmt.Errorf("Unmarshal error:%v, data: %s", err, string(data))
		return nil, err
	}

	if !sr.Success {
		return nil, errors.New(sr.Error)
	}

	return sr.Data, nil
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

type MonthSign struct {
	Sign  string `json:"sign"`
	Array string `json:"arr"`
}

func GetMonthSign(username, year, month string) (*MonthSign, error) {
	url_req := fmt.Sprintf("http://60.209.105.243:8081/ashx/rili.ashx?ty=%s&tm=%s&empnumber=%s", year, month, username)
	resp, err := http.DefaultClient.Get(url_req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ms := &MonthSign{}
	err = json.Unmarshal(data, ms)
	if err != nil {
		return nil, err
	}

	return ms, nil
}
