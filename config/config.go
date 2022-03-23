package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const (
	Path = "config.json"
)

type Config struct {
	ListenAddr string `json:"listen_addr"`
	Mongo      struct {
		Url         string `json:"url"`
		Collections struct {
			User string `json:"user"`
			Task string `json:"task"`
		} `json:"collections"`
	} `json:"mongo"`
	Redis struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
	} `json:"redis"`
	Env string `json:"env"`
}

var globalConfig Config

func init() {
	b, err := ioutil.ReadFile(Path)
	if err != nil {
		log.Printf("No config.json? Use config.json.example to create one.")
		panic(err)
	}
	err = json.Unmarshal(b, &globalConfig)
	if err != nil {
		panic(err)
	}
}

func Get() *Config {
	return &globalConfig
}
