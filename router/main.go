package router

import (
	"loanManagement/handler"
	"loanManagement/instance"
	"loanManagement/repo"
)

type Router interface {
	Fallback() Fallback
	User() User
}

type router struct {
	fallback Fallback
	user     User
}

func Init(instance instance.Instance) Router {
	contextUtil := instance.ContextUtil()
	jwtUtil := instance.JwtUtil()
	passwordUtil := instance.PasswordUtil()
	timeUtil := instance.TimeUtil()

	dbInstance := instance.DatabaseInstance()

	loanRepo := repo.NewLoan(dbInstance)
	scheduledRepaymentRepo := repo.NewScheduledRepayment(dbInstance)
	userRepo := repo.NewUser(dbInstance, passwordUtil)

	userHandler := handler.NewUser(
		loanRepo,
		scheduledRepaymentRepo,
		userRepo,

		dbInstance,

		jwtUtil,
		timeUtil,
	)

	defaultRouter := NewDefault()
	userRouter := NewUser(
		userHandler,

		contextUtil,
		jwtUtil,
		passwordUtil,
	)

	router := router{
		fallback: defaultRouter,
		user:     userRouter,
	}

	return &router
}

func (r *router) Fallback() Fallback {
	return r.fallback
}

func (r *router) User() User {
	return r.user
}
