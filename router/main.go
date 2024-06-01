package router

import (
	"loanManagement/handler"
	"loanManagement/instance"
	"loanManagement/repo"
)

type Router interface {
	Default() Default
	User() User
}

type router struct {
	defaultRouter Default
	userRouter    User
}

func Init(instance instance.Instance) Router {
	jwtUtil := instance.JwtUtil()
	passwordUtil := instance.PasswordUtil()

	dbInstance := instance.DatabaseInstance()

	userRepo := repo.NewUser(dbInstance, passwordUtil)

	userHandler := handler.NewUser(
		userRepo,

		jwtUtil,
	)

	defaultRouter := NewDefault()
	userRouter := NewUser(
		userHandler,

		instance.ContextUtil(),
		instance.JwtUtil(),
		instance.PasswordUtil(),
	)

	router := router{
		defaultRouter: defaultRouter,
		userRouter:    userRouter,
	}

	return &router
}

func (r *router) Default() Default {
	return r.defaultRouter
}

func (r *router) User() User {
	return r.userRouter
}
