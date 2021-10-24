package repositories

import (
	"context"
	"fmt"
	"log"

	character "rickAndMortyApi/internal"

	"go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type characterRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewCharacterRepository() *characterRepository {
	host := "localhost"
	port := 27017

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connections
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Congratulations, you're already connected to MongoDB!")
	collection := client.Database("RickAndMorty").Collection("Characters")

	return &characterRepository{
		client:     client,
		collection: collection,
	}
}

func (r *characterRepository) FetchCharacters() ([]*character.Character, error) {
	var results []*character.Character
	findOptions := options.Find()
	cur, err := r.collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var s character.Character
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
