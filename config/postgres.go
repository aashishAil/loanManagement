package config

import "github.com/spf13/viper"

type postgres struct {
	dbName             string
	host               string
	maxIdleConnections *int
	maxOpenConnections *int
	password           string
	port               int
	sslMode            bool
	user               string
}

func (c *postgres) load() {
	viper.SetEnvPrefix("POSTGRES")
	c.host = viper.GetString("HOST")
	c.port = viper.GetInt("PORT")
	c.user = viper.GetString("USER")
	c.password = viper.GetString("PASSWORD")
	c.dbName = viper.GetString("DB_NAME")
	c.sslMode = viper.GetBool("SSL_MODE")
	idleConnections := viper.GetInt("MAX_IDLE_CONNECTIONS")
	if idleConnections != 0 {
		c.maxIdleConnections = &idleConnections
	}
	openConnections := viper.GetInt("MAX_OPEN_CONNECTIONS")
	if openConnections != 0 {
		c.maxOpenConnections = &openConnections
	}
}
