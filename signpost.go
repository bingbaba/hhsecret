package hhsecret

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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

type SignConfigResp struct {
	Success   bool         `json:"success"`
	Error     string       `json:"error"`
	ErrorCode int          `json:"errorCode"`
	Data      *SingnConfig `json:"data"`
}

type SingnConfig struct {
	ConfigId string `json:"configId"`
}

type SignResp struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error"`
	ErrorCode int        `json:"errorCode"`
	Data      *SingnData `json:"data"`
}

type SingnData struct {
	Id          string   `json:"id"`
	Datetime    int64    `json:"time"`
	Feature     string   `json:"feature"`
	Content     string   `json:"content"`
	PhotoIds    []string `json:"photoIds"`
	ClockInType int      `json:"clockInType"`
	PointId     string   `json:"pointId"`
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
	DateTime int64  `json:"time"`
	Location string `json:"feature"`
	RecordId string `json:"recordId"`
	PointId  string `json:"pointId"`
	//Latitude  float64 `json:"latitude"`
	//Longitude float64 `json:"longitude"`
}

func (s *Sign) GetMinuteSecode() string {
	return time.Unix(s.DateTime/1000, 0).Format("15:04")
}

func (client *Client) signConfigId(lat, lng string) (string, error) {
	configId := "5d007693ebf84b14e8287240_0"

	// configId
	form := url.Values{
		"latitude":  {lat},
		"longitude": {lng},
		"networkId": {"5170b3ede4b0e5e16493be38"},
		"userId":    {client.LoginData.OpenID},
	}
	req, err := client.newHttpReq(
		"POST",
		"/attendance-signapi/config/inner/set/signConfig",
		form)
	if err != nil {
		return configId, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return configId, err
	}
	defer resp.Body.Close()

	scr := &SignConfigResp{}
	err = json.NewDecoder(resp.Body).Decode(scr)
	if err != nil {
		err = fmt.Errorf("Unmarshal error:%v", err)
		return configId, err
	}

	if !scr.Success {
		return configId, errors.New(scr.Error)
	}

	if scr.Data != nil {
		configId = scr.Data.ConfigId
	}

	return configId, nil
}

func (client *Client) Sign() (*SingnData, error) {
	// lng=120.441031
	// bssid=40%3A31%3A3c%3Adf%3Ad7%3A6a
	// configId=5d007693ebf84b14e8287240_18
	// networkId=5170b3ede4b035e16493be37
	// userId=705ac0be-3557-11e7-86f4-44a842003fef
	// ssid=bingbaba
	// lat=36.181592

	// location
	lat := fmt.Sprintf("%0.6f", 36.12998+random.Float64()/100000)
	lng := fmt.Sprintf("%0.6f", 120.41626+random.Float64()/100000)

	configId, err := client.signConfigId(lat, lng)
	if err != nil {
		return nil, err
	}

	// sign in
	form := url.Values{
		"bssid":     {client.LoginInfo.device.GetMac(client.LoginInfo.userName)},
		"configId":  {configId},
		"lat":       {lat},
		"lng":       {lng},
		"userId":    {client.LoginData.OpenID},
		"ssid":      {""},
		"networkId": {client.LoginData.WbNetworkId},
	}
	log.Printf("%s", form.Encode())

	req, err := client.newHttpReq(
		"POST",
		"/attendance-signapi/signservice/sign/signIn",
		form)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	sr := &SignResp{}
	//err = json.NewDecoder(resp.Body).Decode(sr)
	err = json.Unmarshal(data, sr)
	if err != nil {
		err = fmt.Errorf("Unmarshal error:%v", err)
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
	req, err := client.newHttpReq(
		"POST",
		"/attendance-signapi/clockdetail/sign/currentlist",
		form)
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
	url_req := fmt.Sprintf("https://hrit.haier.net:8899/ashx/rili.ashx?ty=%s&tm=%s&empnumber=%s", year, month, username)
	resp, err := http.DefaultClient.Post(url_req, "application/json", nil)
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
