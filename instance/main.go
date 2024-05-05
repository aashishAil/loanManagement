package instance

import (
	"github.com/pkg/errors"
	"loanManagement/util"

	"loanManagement/config"
	dbInstance "loanManagement/database/instance"
	"loanManagement/logger"
)

type Instance interface {
	DatabaseInstance() dbInstance.PostgresDB
	JwtUtil() util.Jwt
	PasswordUtil() util.Password
}

type instance struct {
	dbInstance dbInstance.PostgresDB

	jwtUtil      util.Jwt
	passwordUtil util.Password
}

func Init() (Instance, error) {

	postgresConfig := config.Env.PostgresConfig()
	jwtSigningKey := config.Env.JwtSigningKey()

	databaseInstance, err := dbInstance.NewPostgresDatabase(postgresConfig)
	if err != nil {
		logger.Log.Errorf("failed to create database instance: %v", err)
		return nil, errors.Wrap(err, "failed to create database instance")
	}

	jwtUtil := util.NewJwt(jwtSigningKey)
	passwordUtil := util.NewPassword()

	instance := instance{
		dbInstance: databaseInstance,

		jwtUtil:      jwtUtil,
		passwordUtil: passwordUtil,
	}

	return &instance, nil
}

func (i *instance) DatabaseInstance() dbInstance.PostgresDB {
	return i.dbInstance
}

func (i *instance) JwtUtil() util.Jwt {
	return i.jwtUtil
}

func (i *instance) PasswordUtil() util.Password {
	return i.passwordUtil
}
