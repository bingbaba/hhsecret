package web

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var DefaultCfg *Configure

type Configure struct {
	Port           int      `yaml:"port"`
	LogFile        string   `yaml:"logfile"`
	ConsumerKey    string   `yaml:"consumer_key"`
	ConsumerSecret string   `yaml:"consumer_secret"`
	UserWhiteList  []string `yaml:"UserWhiteList"`
	StaticWebPath  string   `yaml:"staticWebPath"`
}

func init() {
	DefaultCfg = &Configure{}
}

func LoadConfigureByFile(filename string) (err error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(body, DefaultCfg)
	return
}
