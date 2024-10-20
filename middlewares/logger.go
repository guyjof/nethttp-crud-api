package middlewares

import (
	"log"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

// override the WriteHeader method to store the status code
func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

// loggingMiddleware logs the request method, URL, and status code
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().Format(time.UnixDate)

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next(wrapped, r)
		log.Printf("%v %v request from %v to %v responded with %v\n", start, r.Method, r.RemoteAddr, r.URL, wrapped.statusCode)
	}
}
