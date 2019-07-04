package hhsecret

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

var (
	FMT_USERAGENT = "38882/10.1.8(1050);Android 8.0.0;%s;%s;102;%s;deviceId:%s;deviceName:%s %s;clientId:38882;os:Android 8.0.0;brand:%s;model:%s;oem:ihaier;lang:zh-CN;"
	Devices       = []*DeviceInfo{
		{"Huawei", "P30 Pro", "2430*1080", "b4:43:26", ""},
		{"Oppo", "Reno", "2430*1080", "c4:e3:9f", ""},
		{"Vivo", "X27", "2430*1080", "0c:20:d3", ""},
		{"Huawei", "Mate 20 Pro", "3120x1440", "b4:43:26", ""},
		{"Xiaomi", "MIX+2", "1080*2030", "40:31:3c", ""},
		{"Sansung", "Galaxy S10", "3040x1440", "5c:0a:5b", ""},
	}
)

type DeviceInfo struct {
	Brand      string
	Model      string
	Resolution string
	MacPrefix  string
	DeviceId   string
}

func (di *DeviceInfo) UserAgent(devid string) string {
	return fmt.Sprintf(FMT_USERAGENT, di.Brand, di.Model, di.Resolution, devid, di.Brand, di.Model, di.Brand, di.Model)
}

func (di *DeviceInfo) GetMac(username string) string {
	user_hash := fmt.Sprintf("%x", md5.Sum([]byte(username)))
	return fmt.Sprintf("%s:%s:%s:%s", di.MacPrefix, user_hash[0:2], user_hash[2:4], user_hash[4:6])
}

func GetDeviceInfo(username string) DeviceInfo {
	if username == "" {
		return *Devices[0]
	}
	idx, _ := strconv.Atoi(username[1:])

	dev := Devices[idx%len(Devices)]
	dev.DeviceId = fmt.Sprintf("%x", md5.Sum([]byte(username)))
	return *dev
}
