package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Conf Config

func InitConf() {
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Errorf("init viper error. %s", err.Error())
		panic(err)
	}
	err := viper.Unmarshal(&Conf)
	if err != nil {
		panic("unmarshal config file error")
	}
}
