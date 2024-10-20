package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"
)

func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {

				msg := "Caught panic: %v, Stack trace: %s"
				log.Printf(msg, err, string(debug.Stack()))

				er := http.StatusInternalServerError
				http.Error(w, http.StatusText(er), er)
			}
		}()
		next(w, r)
	}
}
