package router

import (
	"loanManagement/handler"
	"loanManagement/instance"
	"loanManagement/repo"
)

type Router interface {
	Admin() Admin
	Fallback() Fallback
	User() User
}

type router struct {
	admin    Admin
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

	adminHandler := handler.NewAdmin(
		loanRepo,
		scheduledRepaymentRepo,

		dbInstance,
	)
	userHandler := handler.NewUser(
		loanRepo,
		scheduledRepaymentRepo,
		userRepo,

		dbInstance,

		jwtUtil,
		timeUtil,
	)

	adminRouter := NewAdmin(
		adminHandler,

		contextUtil,
	)
	fallbackRouter := NewFallback()
	userRouter := NewUser(
		userHandler,

		contextUtil,
		jwtUtil,
		passwordUtil,
	)

	routerI := router{
		admin:    adminRouter,
		fallback: fallbackRouter,
		user:     userRouter,
	}

	return &routerI
}

func (r *router) Admin() Admin {
	return r.admin

}

func (r *router) Fallback() Fallback {
	return r.fallback
}

func (r *router) User() User {
	return r.user
}
