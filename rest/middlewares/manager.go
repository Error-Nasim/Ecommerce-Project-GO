package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddleWares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddleWares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddleWares = append(mngr.globalMiddleWares, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	h := next

	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func (mngr *Manager) WrapMux(next http.Handler, middlewares ...Middleware) http.Handler {
	h := next

	for _, middleware := range mngr.globalMiddleWares {
		h = middleware(h)
	}
	return h
}
