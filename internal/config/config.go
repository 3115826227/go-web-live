package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	config Config
)

type Config struct {
	Log struct {
		LogLevel string `json:"log_level"`
		LogPath  string `json:"log_path"`
	} `json:"log"`
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server"`
	Database struct {
		Type       string `json:"type"`
		DataSource string `json:"data_source"`
	} `json:"database"`

	TokenSecret string
}

func GetConfig() Config {
	return config
}

func readConfig() (err error) {
	var data []byte
	if data, err = ioutil.ReadFile("cmd/res/config.yaml"); err != nil {
		err = errors.New(fmt.Sprintf("failed read config file: %s \n", err))
		return
	}
	if err = yaml.Unmarshal(data, &config); err != nil {
		err = errors.New(fmt.Sprintf("failed unmarshal config file: %s \n", err))
	}
	return
}

func init() {
	if err := readConfig(); err != nil {
		panic(err)
	}
}
