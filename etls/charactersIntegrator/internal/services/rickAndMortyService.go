package services

import (
	dtos "charactersIntegrator/internal/dtos"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RickAndMortyService interface {
	GetCharacters(page int) (characters *dtos.AllCharactersDto, err error)
}

type rickAndMortyService struct {
	url string
}

func NewRickAndMortyService() *rickAndMortyService {
	return &rickAndMortyService{
		url: "https://rickandmortyapi.com/api/character?page=%d",
	}
}

func (r *rickAndMortyService) GetCharacters(page int) (characters *dtos.AllCharactersDto, err error) {
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
