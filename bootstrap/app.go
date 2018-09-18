package bootstrap

import (
	"github.com/spf13/viper"
)

func BootstrapApp() {
	initViper()
}

func initViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		panic("Cannot load config file")
	}
}
