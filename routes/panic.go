package routes

import "net/http"

// this function will panic
func panic(w http.ResponseWriter, r *http.Request) {
	var tmp *int
	*tmp += 1
}
