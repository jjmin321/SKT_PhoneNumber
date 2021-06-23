package config

import (
	"SKT_PhoneNumber/lib/log"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

var Config config

type config struct {
	Skt skt
}

type skt struct {
	URL string `toml:"url"`
	ViewId string `toml:"viewId"`
	ServiceId string `toml:"serviceId"`
	Chg_cd string `toml:"chg_cd"`
	Co_cd string `toml:"co_cd"`
	Line_num_from string `toml:"line_num_from"`
	Line_num_to string `toml:"line_num_to"`
}

func init() {
	configBytes, err := ioutil.ReadFile("config/config.toml")
	if err != nil {
		panic(err)
	}

	_, err = toml.Decode(string(configBytes), &Config)
	if err != nil {
		panic(err)
	}

	log.Info.Print("[CONF] 환경설정 초기화")
}