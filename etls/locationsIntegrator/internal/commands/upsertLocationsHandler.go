package commands

import (
	"fmt"
	entities "locationsIntegrator/internal"
	dtos "locationsIntegrator/internal/dtos"
	repositories "locationsIntegrator/internal/repositories"
	services "locationsIntegrator/internal/services"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type upsertLocationsHandler struct {
	service    services.RickAndMortyService
	repository repositories.RickAndMortyRepository
}

func NewUpsertLocationsHandler() *upsertLocationsHandler {
	return &upsertLocationsHandler{
		service:    services.NewRickAndMortyService(),
		repository: repositories.NewRickAndMortyRepository(),
	}
}

func (h *upsertLocationsHandler) Handler() (err error) {
	nextPage := true

	upsertResults := &mongo.UpdateResult{
		MatchedCount:  0,
		ModifiedCount: 0,
		UpsertedCount: 0,
	}

	for i := 1; nextPage; i++ {
		result, err := h.service.GetLocations(i)

		if err != nil {
			log.Fatalln(err)
		}

		if result != nil && result.Results != nil {
			for i := 0; i < len(result.Results); i++ {
				location := MapToEntity(result.Results[i])

				upsertResult, upsertError := h.repository.Upsert(location.Id, location)

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

func MapToEntity(locationDto dtos.LocationDto) entities.Location {
	location := entities.Location{
		Entity:    entities.Entity{Id: locationDto.ID},
		Name:      locationDto.Name,
		Type:      locationDto.Type,
		Dimension: locationDto.Dimension,
		Residents: locationDto.Residents,
		URL:       locationDto.URL,
		Created:   locationDto.Created,
	}

	return location
}
