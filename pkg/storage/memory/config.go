package memory

import (
	"flag"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var cfg Config

const (
	DefaultPresicion         = 30 * time.Second
	DefaultPartitionDuration = 2 * time.Hour
)

type Config struct {
	// usually 15s | 30s | 1m | 2m | 5m
	Precision time.Duration `yaml:"precision"`

	// define the size of a block(e.g. 2 hours)
	PartitionDuration time.Duration `yaml:"partitionduration"`
}

func (c *Config) RegisterFlag() {
	flag.DurationVar(&cfg.Precision, "storage.memory.precision", DefaultPresicion, "define the precision of data points")
	flag.DurationVar(&cfg.PartitionDuration, "storage.memory.partitionduration", DefaultPartitionDuration, "define the size of a timeseries block")
}

// func init() {
// 	SetDefaultConfig()
// 	ParseConfig()
// 	logrus.Infof("memory storage config: %+v", cfg)
// }

func LoadConfig() {
	SetDefaultConfig()
	ParseConfig()
	logrus.Infof("memory storage config: %+v", cfg)
}

func SetDefaultConfig() {
	viper.SetDefault("storage.memory.precision", DefaultPresicion)
	viper.SetDefault("storage.memory.partitionduration", DefaultPartitionDuration)
}

func ParseConfig() {
	cfg.Precision = viper.GetDuration("storage.memory.precision")
	cfg.Precision = viper.GetDuration("storage.memory.partitionduration")
}
