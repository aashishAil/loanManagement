package instance

import (
	"fmt"

	"loanManagement/logger"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type PostgresDB interface {
	GetReadableDb() *gorm.DB
	GetWritableDb() *gorm.DB
	GetTransactionDb() (PostgresTransactionDB, error)
}

type postgresDB struct {
	connectionUrl      string
	maxIdleConnections *int
	maxOpenConnections *int
	database           *gorm.DB
}

func (d *postgresDB) initialize() error {
	db, err := gorm.Open(postgres.Open(d.connectionUrl), &gorm.Config{
		SkipDefaultTransaction: true, // writes should not be run in a default transaction
		Logger:                 gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		return errors.Wrap(err, "failed to connect to database")
	}

	if d.maxIdleConnections != nil || d.maxOpenConnections != nil {
		config := dbresolver.Register(dbresolver.Config{}) // use this for registering read/write connections
		if d.maxIdleConnections != nil {
			config = config.SetMaxIdleConns(*d.maxIdleConnections)
		}

		if d.maxOpenConnections != nil {
			config = config.SetMaxOpenConns(*d.maxOpenConnections)
		}
		err := db.Use(config)
		if err != nil {
			return errors.Wrap(err, "failed to add config to database")
		}
	}

	d.database = db

	return nil
}

func (d *postgresDB) GetReadableDb() *gorm.DB {
	return d.database
}

func (d *postgresDB) GetWritableDb() *gorm.DB {
	d.database.Begin()
	return d.database
}

func (d *postgresDB) GetTransactionDb() (PostgresTransactionDB, error) {
	txnDb := d.GetWritableDb().Begin()
	localTransactionI := NewPostgresTransactionDB(txnDb)

	err := localTransactionI.CheckError()
	if err != nil {
		logger.Log.Error("unable to start transaction", logger.Error(err))
		return nil, err
	}

	return localTransactionI, nil
}

func NewPostgresDatabase(config PostgresDbConfig) (PostgresDB, error) {
	connectionUrl := ""
	if config.Host != "" {
		connectionUrl += fmt.Sprintf("host=%s ", config.Host)
	}

	if config.User != "" {
		connectionUrl += fmt.Sprintf("user=%s ", config.User)
	}

	if config.Password != "" {
		connectionUrl += fmt.Sprintf("password=%s ", config.Password)
	}

	if config.DbName != "" {
		connectionUrl += fmt.Sprintf("dbname=%s ", config.DbName)
	}

	if config.Port != 0 {
		connectionUrl += fmt.Sprintf("port=%d ", config.Port)
	}

	if config.SslMode != "" {
		connectionUrl += fmt.Sprintf("sslmode=%s ", config.SslMode)
	}
	db := postgresDB{
		connectionUrl:      connectionUrl,
		maxIdleConnections: config.MaxIdleConnections,
		maxOpenConnections: config.MaxOpenConnections,
	}

	err := db.initialize()
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize postgres database")
	}

	return &db, nil
}
