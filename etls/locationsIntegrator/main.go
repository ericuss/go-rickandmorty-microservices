package main

import (
	entities "locationsIntegrator/internal"
	repositories "locationsIntegrator/internal/repositories"
	services "locationsIntegrator/internal/services"
	"log"
)

func main() {
	log.Println("Starting app...")

	service := services.NewRickAndMortyService()
	repository := repositories.NewRickAndMortyRepository()
	result, getLocationsError := service.GetLocations()

	if getLocationsError != nil {
		log.Fatalln(getLocationsError)
	}

	if result != nil && result.Results != nil {
		for i := 0; i < len(result.Results); i++ {
			locationDto := result.Results[i]

			log.Println(result.Results[i].Name)

			location := entities.Location{
				Id:        locationDto.ID,
				Name:      locationDto.Name,
				Type:      locationDto.Type,
				Dimension: locationDto.Dimension,
				Residents: locationDto.Residents,
				URL:       locationDto.URL,
				Created:   locationDto.Created,
			}

			id, addError := repository.Add(location)

			if addError != nil {
				log.Fatalln(addError)
			}

			log.Println(id)
		}
	}

	log.Println("Finished")
}
