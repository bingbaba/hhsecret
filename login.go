package hhsecret

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	URL_LOGIN     = "http://hdxtapp.haier.net/openaccess/user/login"
	FMT_LOGININFO = `{"eid":"","userName":"%s","password":"%s","appClientId":"10201","deviceId":"%s","deviceType":"%s","ua":"%s"}`
)

type LoginResp struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error"`
	ErrorCode int        `json:"errorCode"`
	Data      *LoginData `json:"data"`
}

type LoginData struct {
	Birthday         string   `json:"birthday"`
	Phone            string   `json:"phone"`
	FullPinyinCode   string   `json:"fullPinyinCode"`
	LastUpdateTime   string   `json:"lastUpdateTime"`
	Department       string   `json:"department"`
	BindedPhone      string   `json:"bindedPhone"`
	OauthTokenSecret string   `json:"oauth_token_secret"`
	Eid              string   `json:"eid"`
	Phones           string   `json:"phones"`
	UserType         int      `json:"userType"`
	Id               string   `json:"id"`
	OrgId            string   `json:"orgId"`
	OrgLongName      string   `json:"orgLongName"`
	Token            string   `json:"token"`
	OauthToken       string   `json:"oauth_token"`
	Name             string   `json:"name"`
	WbUserId         string   `json:"wbUserId"`
	Gender           int      `json:"gender"`
	FirstPinyin      string   `json:"firstPinyin"`
	OfficePhone2     string   `json:"officePhone2"`
	JobTitle         string   `json:"jobTitle"`
	FullPinyin       string   `json:"fullPinyin"`
	Status           int      `json:"status"`
	OfficePhone1     string   `json:"officePhone1"`
	OrgInfoId        string   `json:"orgInfoId"`
	FirstPinyinCode  string   `json:"firstPinyinCode"`
	Emails           string   `json:"emails"`
	CompanyName      string   `json:"companyName"`
	BindedEmail      string   `json:"bindedEmail"`
	Weights          int      `json:"weights"`
	Email            string   `json:"email"`
	IsHidePhone      int      `json:"isHidePhone"`
	IsAdmin          int      `json:"isAdmin"`
	OId              string   `json:"oId"`
	PhotoUrl         string   `json:"photoUrl"`
	WbNetworkId      string   `json:"wbNetworkId"`
	Forbidden        []string `json:"forbidden"`
	OrgUserType      int      `json:"orgUserType"`
	OpenID           string   `json:"openId"`
}

type LoginInfo struct {
	userName  string
	password  string
	devid     string
	devtype   string
	useragent string
}

func NewLoginInfo(username, password string) *LoginInfo {
	password_encrypt, err := DesEncrypt([]byte(password), []byte(username)[0:8])
	if err != nil {
		panic(err)
	}

	password_encrypt_str := base64.StdEncoding.EncodeToString(password_encrypt)
	l := &LoginInfo{
		userName:  username,
		password:  password_encrypt_str,
		devtype:   "MHA-AL00 Build/HUAWEIMHA-AL00",
		useragent: "Linux; Android 7.0; MHA-AL00 Build/HUAWEIMHA-AL00; wv",
	}
	l.setDevID()
	return l
}

func (l *LoginInfo) Do() (*LoginData, error) {
	req, err := http.NewRequest("POST", URL_LOGIN, bytes.NewReader([]byte(l.ToString())))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", l.useragent)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ld := &LoginResp{}
	err = json.Unmarshal(data, ld)
	if err != nil {
		return nil, err
	}

	if ld.ErrorCode != 100 {
		return nil, errors.New(ld.Error)
	}
	return ld.Data, nil
}

func (l *LoginInfo) ToString() string {
	return fmt.Sprintf(FMT_LOGININFO,
		l.userName, l.password,
		l.devid, l.devtype,
		l.useragent)
}

func (l *LoginInfo) setDevID() {
	l.devid = fmt.Sprintf("%X", md5.Sum([]byte(l.userName)))
}
