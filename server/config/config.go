package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var conf *viper.Viper

func init() {
	conf = viper.New()
	conf.SetConfigName("config")
	conf.SetConfigType("yaml")
	conf.AddConfigPath(".")
	if err := conf.ReadInConfig(); err != nil {
		logrus.Println("加载配置文件失败:", err)
		os.Exit(1)
	}
}

func GetConfig() *viper.Viper {
	return conf
}
