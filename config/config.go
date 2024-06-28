package config

import (
	"github.com/RINOHeinrich/multiserviceauth/helper"
)

var Config helper.AppConfig

func LoadConfig(filename string) {
	err := Config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

}
