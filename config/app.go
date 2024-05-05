package config

import "github.com/spf13/viper"

type app struct {
	isDevelopment bool
	name          string
	port          string
}

func (a *app) load() {
	viper.SetEnvPrefix("APP")
	a.name = viper.GetString("NAME")
	a.isDevelopment = viper.GetBool("IS_DEVELOPMENT")
	a.port = viper.GetString("PORT")
}
