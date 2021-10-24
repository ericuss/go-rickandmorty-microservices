package server

import (
	"encoding/json"
	"net/http"

	character "rickAndMortyApi/internal"

	"github.com/gorilla/mux"
)

type charactersController struct {
	router     *mux.Router
	repository character.CharacterRepository
}

func NewCharactersController(repository character.CharacterRepository, r *mux.Router) Server {
	a := &charactersController{repository: repository}

	r.HandleFunc("/api/characters", a.fetch).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *charactersController) routes() {
	a.router.HandleFunc("/api/characters", a.fetch).Methods(http.MethodGet)
}
func (a *charactersController) Router() mux.Router {
	return *a.router
}

func (a *charactersController) fetch(w http.ResponseWriter, r *http.Request) {
	beers, _ := a.repository.FetchCharacters()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(beers)
}
