package http

import (
	"flag"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var cfg Config

type Config struct {
	Port int `yaml:"port"`
}

func (c *Config) RegisterFlag() {
	flag.IntVar(&cfg.Port, "http.port", 9009, "define the precision of data points")
}

// func init() {
// 	SetDefaultConfig()
// 	ParseConfig()
// 	logrus.Infof("http config: %+v", cfg)
// }

func LoadConfig() {
	SetDefaultConfig()
	ParseConfig()
	logrus.Infof("http config: %+v", cfg)
}

func SetDefaultConfig() {
	viper.SetDefault("http.port", 9009)
}

func ParseConfig() {
	viper.GetInt("http.port")
}
