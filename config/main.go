package config

import (
	_ "github.com/joho/godotenv/autoload" // auto-load the env variables from the .env file
	"github.com/spf13/viper"

	"loanManagement/database/instance"
)

type Config interface {
	PostgresConfig() instance.PostgresDbConfig
	ServerPort() string
	AppName() string
	IsDevelopment() bool
}
type configType struct {
	app      app
	postgres postgres
}

var Env Config

func (c *configType) load() {
	viper.AutomaticEnv()
	c.app = app{}
	c.app.load()
	c.postgres = postgres{}
	c.postgres.load()

}

func (c *configType) PostgresConfig() instance.PostgresDbConfig {
	return instance.PostgresDbConfig{
		Host:               c.postgres.host,
		Port:               c.postgres.port,
		User:               c.postgres.user,
		Password:           c.postgres.password,
		DbName:             c.postgres.dbName,
		SslMode:            c.postgres.sslMode,
		MaxIdleConnections: c.postgres.maxIdleConnections,
		MaxOpenConnections: c.postgres.maxOpenConnections,
	}
}

func (c *configType) ServerPort() string {
	return "localhost:" + c.app.port
}

func (c *configType) AppName() string {
	return c.app.name
}

func (c *configType) IsDevelopment() bool {
	return c.app.isDevelopment
}

func init() {
	config := configType{}
	config.load()
	Env = &config
}
