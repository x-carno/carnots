package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/x-carno/carnots/pkg/storage/memory"
	"github.com/x-carno/carnots/pkg/svc/http"
)

type Data struct {
	Value     int16
	ID        uint32
	Timestamp uint64
}

type Flag struct {
	Body [4]byte
}

var configFilePath string

func main() {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)

	// cfg.RegisterFlag()
	flag.StringVar(&configFilePath, "config.file", "config.yaml", "config file path")
	flag.Parse()

	loadConfig()

	server := http.NewServer()
	server.ListenHttp()
}

func loadConfig() {
	viper.SetConfigFile(configFilePath)
	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorln("load config file failed : ", err.Error())
		os.Exit(1)
	}
	loadModuleConfigs()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		loadModuleConfigs()
	})
	viper.WatchConfig()

	// buf, err := ioutil.ReadFile(configFilePath)
	// if err != nil {
	// 	logrus.Errorln("load config file failed : ", err.Error())
	// }
	// err = yaml.Unmarshal(buf, &cfg)
	// if err != nil {
	// 	logrus.Errorln("unmarshall failed : ", err.Error())
	// }
	// logrus.Debugln("config : ", cfg)
}

func loadModuleConfigs() {
	http.LoadConfig()
	memory.LoadConfig()
}

func writeInt64Bits(i int64, nbits uint) error {
	var u uint64
	if i >= 0 || nbits >= 64 {
		u = uint64(i)
	} else {
		u = uint64(1<<nbits + i)
	}
	fmt.Println(u)
	return nil
}
