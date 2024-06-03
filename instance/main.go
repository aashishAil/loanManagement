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
	ContextUtil() util.Context
	JwtUtil() util.Jwt
	PasswordUtil() util.Password
	TimeUtil() util.Time
}

type instance struct {
	dbInstance dbInstance.PostgresDB

	contextUtil  util.Context
	jwtUtil      util.Jwt
	passwordUtil util.Password
	timeUtil     util.Time
}

func Init() (Instance, error) {
	logger.Init(config.Env.IsDevelopment())
	postgresConfig := config.Env.PostgresConfig()
	jwtSigningKey := config.Env.JwtSigningKey()

	databaseInstance, err := dbInstance.NewPostgresDatabase(postgresConfig)
	if err != nil {
		logger.Log.Errorf("failed to create database instance: %v", err)
		return nil, errors.Wrap(err, "failed to create database instance")
	}

	contextUtil := util.NewContext()
	jwtUtil := util.NewJwt(jwtSigningKey)
	passwordUtil := util.NewPassword()
	timeUtil := util.NewTime()

	instance := instance{
		dbInstance: databaseInstance,

		contextUtil:  contextUtil,
		jwtUtil:      jwtUtil,
		passwordUtil: passwordUtil,
		timeUtil:     timeUtil,
	}

	return &instance, nil
}

func (i *instance) DatabaseInstance() dbInstance.PostgresDB {
	return i.dbInstance
}

func (i *instance) ContextUtil() util.Context {
	return i.contextUtil
}

func (i *instance) JwtUtil() util.Jwt {
	return i.jwtUtil
}

func (i *instance) PasswordUtil() util.Password {
	return i.passwordUtil
}

func (i *instance) TimeUtil() util.Time {
	return i.timeUtil
}
