package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var GlobeCfg *Config

type Config struct {
	Mysql SqlConfig `yaml:"mysql"`
}

type SqlConfig struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pwd  string `yaml:"pwd"`
	DB   string `yaml:"db"`
}

func LoadConfig() {
	var cfg Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.ReadInConfig()
	err := viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Println("config error", err)
		return
	}
	GlobeCfg = &cfg
}
