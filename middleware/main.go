package middleware

type Middleware interface {
	Server() Server
}

type middleware struct {
	server Server
}

func Init() Middleware {
	instance := middleware{
		server: NewServer(),
	}
	return &instance
}

func (m *middleware) Server() Server {
	return m.server
}
