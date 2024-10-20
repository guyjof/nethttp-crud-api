package middlewares

import (
	"fmt"
	"net/http"
)

func authenticateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer my_token" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized")
			return
		}
		next(w, r)
	}
}
