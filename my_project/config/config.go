package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"my_project/logger"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type"`
		BindIp string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {

	once.Do(func() {
		Log := logger.GetLogger()
		Log.Info("read config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			Log.Info(help)
			//Log.Fatal(err)
		}
	})
	return instance
}
