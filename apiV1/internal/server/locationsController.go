package server

import (
	"encoding/json"
	"net/http"

	location "rickAndMortyApi/internal"

	"github.com/gorilla/mux"
)

type locationsController struct {
	router     *mux.Router
	repository location.LocationRepository
}

func NewLocationsController(repository location.LocationRepository, r *mux.Router) Server {
	a := &locationsController{repository: repository}

	r.HandleFunc("/api/locations", a.fetch).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *locationsController) routes() {
	a.router.HandleFunc("/api/locations", a.fetch).Methods(http.MethodGet)
}
func (a *locationsController) Router() mux.Router {
	return *a.router
}

func (a *locationsController) fetch(w http.ResponseWriter, r *http.Request) {
	beers, _ := a.repository.Fetch()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(beers)
}
