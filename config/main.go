package config

import (
	"loanManagement/database/instance"
)

type Config interface {
	PostgresConfig() instance.PostgresDbConfig
}
type config struct {
	postgres postgresConfig
}

func (c *config) load() {
	c.postgres = postgresConfig{}
	c.postgres.load()
}

func (c *config) PostgresConfig() instance.PostgresDbConfig {
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

func NewConfig() Config {
	config := &config{}
	config.load()
	return config
}
