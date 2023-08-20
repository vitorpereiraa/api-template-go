package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type settings struct {
	SERVER_PORT string `mapstructure:"SERVER_PORT"`
	SERVER_HOST string `mapstructure:"SERVER_HOST"`
}

var Settings settings

func loadSettings() {
	vp := viper.New()

	vp.SetConfigFile(".env")
	vp.AddConfigPath(".")

	vp.AutomaticEnv()

	if err := vp.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error loading the settings file: %w", err))
	}

	if err := vp.Unmarshal(&Settings); err != nil {
		panic(fmt.Errorf("fatal error parsing the settings file: %w", err))
	}
}
