package instance

type PostgresDbConfig struct {
	Host               string
	Port               int
	User               string
	Password           string
	DbName             string
	SslMode            bool
	MaxIdleConnections *int
	MaxOpenConnections *int
}
