package main

import (
	commands "charactersIntegrator/internal/commands"
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

	upsert := commands.NewUpsertCharactersHandler()
	upsert.Handler()

	log.Println("Finished")
}
