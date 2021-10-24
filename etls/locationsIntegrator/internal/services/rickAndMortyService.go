package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	dtos "locationsIntegrator/internal/dtos"
	"net/http"
)

type rickAndMortyService struct {
}

func NewRickAndMortyService() *rickAndMortyService {
	return &rickAndMortyService{}
}

func (r *rickAndMortyService) GetLocations() (locations *dtos.AllLocationsDto, err error) {
	response, err := http.Get(fmt.Sprintf("%v", "https://rickandmortyapi.com/api/location?page=1"))
	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &locations)
	if err != nil {
		return nil, err
	}
	return
}
