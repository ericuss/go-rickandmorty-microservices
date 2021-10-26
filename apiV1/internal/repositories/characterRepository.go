package repositories

import (
	"context"
	"log"

	entities "rickAndMortyApi/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CharacterRepository interface {
	RepositoryBase
	Fetch() ([]*entities.Character, error)
}

type characterRepository struct {
	repositoryBase
}

func NewCharacterRepository() *characterRepository {
	return &characterRepository{
		repositoryBase: *NewRepositoryBase("Characters"),
	}
}

func (r *characterRepository) Fetch() ([]*entities.Character, error) {
	var results []*entities.Character
	findOptions := options.Find()
	cur, err := r.repositoryBase.collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var s entities.Character
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &s)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return results, nil
}
