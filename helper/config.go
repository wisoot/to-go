package helper

import (
	"github.com/spf13/viper"
)

func Config(name string) string {
	return viper.Get(name).(string)
}
