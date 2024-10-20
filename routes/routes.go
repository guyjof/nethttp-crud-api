package routes

import (
	"net/http"

	"github.com/guyjof/nethttp-crud-api/middlewares"
)

func RegisterRoutes() {
	http.HandleFunc("GET /pokemon", getPokemon)          //get all pokemons
	http.HandleFunc("GET /pokemon/{id}", getPokemonByID) // get a pokemon by id

	ah := createPokemon // this wraps the createPokemon handler with the middlewares
	for _, m := range middlewares.Middleware {
		ah = m(ah)
	}
	http.HandleFunc("POST /pokemon", ah) // create a pokemon only if authenticated

	http.HandleFunc("/panic", middlewares.RecoveryMiddleware(panic)) // this route will panic and the middleware will recover
}
