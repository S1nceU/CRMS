package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var Val Config

type DatabaseConfig struct {
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
	Network  string `mapstructure:"NETWORK"`
	Server   string `mapstructure:"SERVER"`
	Port     int    `mapstructure:"PORT"`
	Database string `mapstructure:"DATABASE"`
}

type Config struct {
	Mode            string `mapstructure:"MODE"`
	Port            int    `mapstructure:"PORT"`
	FrontendOrigin  string `mapstructure:"FRONTEND_ORIGIN"`
	CookieSecure    bool   `mapstructure:"COOKIE_SECURE"`
	*DatabaseConfig `mapstructure:"DATABASE"`
}

// Init is a function to read config.yaml
func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := viper.Unmarshal(&Val)
		if err != nil {
			panic(fmt.Errorf("unable to decode into struct, %v", err))
		}
	})

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %v ", err))
	}

	if err := viper.Unmarshal(&Val); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	log.Println("Read config successfully")
}
