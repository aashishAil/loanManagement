package instance

import "gorm.io/gorm"

type PostgresTransactionDB interface {
	CheckError() error
	Commit() error
	Get() *gorm.DB
	Rollback()
}

type postgresTransactionDB struct {
	txnDB *gorm.DB
}

func (t *postgresTransactionDB) CheckError() error {
	return t.txnDB.Error
}

func (t *postgresTransactionDB) Commit() error {
	return t.txnDB.Commit().Error
}

func (t *postgresTransactionDB) Get() *gorm.DB {
	return t.txnDB
}

func (t *postgresTransactionDB) Rollback() {
	t.txnDB.Rollback()
}

func NewPostgresTransactionDB(txnDB *gorm.DB) PostgresTransactionDB {
	return &postgresTransactionDB{
		txnDB: txnDB,
	}
}
