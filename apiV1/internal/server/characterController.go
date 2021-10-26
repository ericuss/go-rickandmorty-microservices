package server

import (
	"encoding/json"
	"net/http"

	repositories "rickAndMortyApi/internal/repositories"

	"github.com/gorilla/mux"
)

type charactersController struct {
	router     *mux.Router
	repository repositories.CharacterRepository
}

func NewCharactersController(repository repositories.CharacterRepository, r *mux.Router) Server {
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
	beers, _ := a.repository.Fetch()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(beers)
}
