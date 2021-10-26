package repositories

import (
	"context"
	"log"

	entities "rickAndMortyApi/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LocationRepository interface {
	RepositoryBase
	Fetch() ([]*entities.Location, error)
}

type locationRepository struct {
	repositoryBase
}

func NewLocationRepository() *locationRepository {
	return &locationRepository{
		repositoryBase: *NewRepositoryBase("Locations"),
	}
}

func (r *locationRepository) Fetch() ([]*entities.Location, error) {
	var results []*entities.Location
	findOptions := options.Find()
	cur, err := r.repositoryBase.collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var s entities.Location
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
