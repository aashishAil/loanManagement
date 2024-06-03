package instance

import "gorm.io/gorm"

type PostgresDbConfig struct {
	Host               string
	Port               int
	User               string
	Password           string
	DbName             string
	SslMode            string
	MaxIdleConnections *int
	MaxOpenConnections *int
	EnableDebugMode    bool
}

var ErrNoRecordFound = gorm.ErrRecordNotFound
