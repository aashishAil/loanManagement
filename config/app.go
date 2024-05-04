package config

import "github.com/spf13/viper"

type app struct {
	name          string
	isDevelopment bool
}

func (a *app) load() {
	viper.SetEnvPrefix("APP")
	a.name = viper.GetString("NAME")
	a.isDevelopment = viper.GetBool("IS_DEVELOPMENT")
}
