package services

import (
	dtos "charactersIntegrator/internal/dtos"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type rickAndMortyService struct {
}

func NewRickAndMortyService() *rickAndMortyService {
	return &rickAndMortyService{}
}

func (r *rickAndMortyService) GetCharacters() (characters *dtos.AllCharactersDto, err error) {
	response, err := http.Get(fmt.Sprintf("%v", "https://charactersIntegrator.com/api/character?page=1"))
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
