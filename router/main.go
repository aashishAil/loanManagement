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
