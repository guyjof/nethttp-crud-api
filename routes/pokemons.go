package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/guyjof/nethttp-crud-api/db"
)

func getPokemon(w http.ResponseWriter, r *http.Request) {
	// Marshal the PokemonDB into JSON
	jsonData, err := json.Marshal(db.PokemonDB)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func getPokemonByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	id := r.PathValue("id")

	for _, pokemon := range db.PokemonDB {
		if id == fmt.Sprintf("%d", pokemon.ID) {
			fmt.Fprintf(w, "Pokemon #%v is %v\n", pokemon.ID, pokemon.Name)
			return
		}
	}
}

func createPokemon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Pokemon created\n")
}
