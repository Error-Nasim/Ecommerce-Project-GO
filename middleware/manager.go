package middleware

import "net/http"

type middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddleWares []middleware
}
