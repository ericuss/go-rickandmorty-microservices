package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	dtos "locationsIntegrator/internal/dtos"
	"net/http"
)

type RickAndMortyService interface {
	GetLocations(page int) (locations *dtos.AllLocationsDto, err error)
}

type rickAndMortyService struct {
	url string
}

func NewRickAndMortyService() *rickAndMortyService {
	return &rickAndMortyService{
		url: "https://rickandmortyapi.com/api/location?page=%d",
	}
}

func (r *rickAndMortyService) GetLocations(page int) (characters *dtos.AllLocationsDto, err error) {
	response, err := http.Get(fmt.Sprintf(r.url, page))
	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &characters)
	if err != nil {
		return nil, err
	}
	return
}
