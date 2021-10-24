package main

import (
	"log"
	repositories "rickAndMortyApi/internal/repositories"
)

func main() {
	log.Println("Starting app...")

	characterRepository := repositories.NewCharacterRepository()

	characters, err := characterRepository.FetchCharacters()

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(characters); i++ {
		log.Println(characters[i].Name)
	}
	log.Println("Finished")
}
