package commands

import (
	entities "charactersIntegrator/internal"
	dtos "charactersIntegrator/internal/dtos"
	repositories "charactersIntegrator/internal/repositories"
	services "charactersIntegrator/internal/services"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type upsertCharactersHandler struct {
	service    services.RickAndMortyService
	repository repositories.RickAndMortyRepository
}

func NewUpsertCharactersHandler() *upsertCharactersHandler {
	return &upsertCharactersHandler{
		service:    services.NewRickAndMortyService(),
		repository: repositories.NewRickAndMortyRepository(),
	}
}

func (h *upsertCharactersHandler) Handler() (err error) {
	nextPage := true

	upsertResults := &mongo.UpdateResult{
		MatchedCount:  0,
		ModifiedCount: 0,
		UpsertedCount: 0,
	}

	for i := 1; nextPage; i++ {
		result, err := h.service.GetCharacters(i)

		if err != nil {
			log.Fatalln(err)
		}

		if result != nil && result.Results != nil {
			for i := 0; i < len(result.Results); i++ {
				character := MapToEntity(result.Results[i])

				upsertResult, upsertError := h.repository.Upsert(character.Id, character)

				if upsertError != nil {
					log.Fatalln(upsertError)
				}

				upsertResults.MatchedCount += upsertResult.MatchedCount
				upsertResults.ModifiedCount += upsertResult.ModifiedCount
				upsertResults.UpsertedCount += upsertResult.UpsertedCount
			}
		}

		nextPage = len(result.Info.Next) > 0
	}

	log.Println(fmt.Printf("Upsert Matched: %d ", upsertResults.MatchedCount))
	log.Println(fmt.Printf("Upsert Modified: %d ", upsertResults.ModifiedCount))
	log.Println(fmt.Printf("Upsert Upserted: %d ", upsertResults.UpsertedCount))

	return
}

func MapToEntity(characterDto dtos.CharacterDto) entities.Character {
	character := entities.Character{
		Entity:  entities.Entity{Id: characterDto.ID},
		Name:    characterDto.Name,
		Status:  characterDto.Status,
		Species: characterDto.Species,
		Type:    characterDto.Type,
		Gender:  characterDto.Gender,
		Image:   characterDto.Image,
		Episode: characterDto.Episode,
		URL:     characterDto.URL,
		Created: characterDto.Created,
	}

	character.Origin.Name = characterDto.Origin.Name
	character.Origin.URL = characterDto.Origin.URL
	character.Location.Name = characterDto.Location.Name
	character.Location.URL = characterDto.Location.URL

	return character
}
