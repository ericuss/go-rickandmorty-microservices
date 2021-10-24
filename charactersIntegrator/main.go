package main

import (
	entities "charactersIntegrator/internal"
	repositories "charactersIntegrator/internal/repositories"
	services "charactersIntegrator/internal/services"
	"log"
)

func main() {
	log.Println("Starting app...")

	service := services.NewRickAndMortyService()
	repository := repositories.NewRickAndMortyRepository()
	result, getCharacterError := service.GetCharacters()

	if getCharacterError != nil {
		log.Fatalln(getCharacterError)
	}

	if result != nil && result.Results != nil {
		for i := 0; i < len(result.Results); i++ {
			characterDto := result.Results[i]

			log.Println(result.Results[i].Name)

			character := entities.Character{
				Id:      characterDto.ID,
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

			id, addError := repository.Add(character)

			if addError != nil {
				log.Fatalln(addError)
			}

			log.Println(id)
		}
	}

	log.Println("Finished")
}
