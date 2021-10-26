package main

import (
	commands "locationsIntegrator/internal/commands"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {

	// try t load .env file
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Error loading .env")
	}
}

func main() {
	log.Println("Starting app...")

	upsert := commands.NewUpsertLocationsHandler()
	upsert.Handler()
	// service := services.NewRickAndMortyService()
	// repository := repositories.NewRickAndMortyRepository()
	// result, getLocationsError := service.GetLocations()

	// if getLocationsError != nil {
	// 	log.Fatalln(getLocationsError)
	// }

	// if result != nil && result.Results != nil {
	// 	for i := 0; i < len(result.Results); i++ {
	// 		locationDto := result.Results[i]

	// 		log.Println(result.Results[i].Name)

	// 		location := entities.Location{
	// 			Id:        locationDto.ID,
	// 			Name:      locationDto.Name,
	// 			Type:      locationDto.Type,
	// 			Dimension: locationDto.Dimension,
	// 			Residents: locationDto.Residents,
	// 			URL:       locationDto.URL,
	// 			Created:   locationDto.Created,
	// 		}

	// 		id, addError := repository.Add(location)

	// 		if addError != nil {
	// 			log.Fatalln(addError)
	// 		}

	// 		log.Println(id)
	// 	}
	// }

	log.Println("Finished")
}
