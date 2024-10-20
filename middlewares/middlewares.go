package middlewares

import "net/http"

var Middleware = []func(http.HandlerFunc) http.HandlerFunc{
	authenticateMiddleware,
	loggingMiddleware,
}
