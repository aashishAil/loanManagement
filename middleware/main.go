package middleware

import "loanManagement/util"

type Middleware interface {
	Auth() Auth
	Server() Server
}

type middleware struct {
	auth   Auth
	server Server
}

func Init(contextUtil util.Context, jwtUtil util.Jwt) Middleware {
	instance := middleware{
		auth:   NewAuth(contextUtil, jwtUtil),
		server: NewServer(),
	}
	return &instance
}

func (m *middleware) Auth() Auth {
	return m.auth
}

func (m *middleware) Server() Server {
	return m.server
}
