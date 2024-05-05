package router

type Router interface {
	Default() Default
}

type router struct {
	defaultRouter Default
}

func Init() Router {
	defaultRouter := NewDefault()

	router := router{
		defaultRouter: defaultRouter,
	}

	return &router
}

func (r *router) Default() Default {
	return r.defaultRouter
}
