package Repositorys

import (
	entities "charactersIntegrator/internal"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type rickAndMortyRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewRickAndMortyRepository() *rickAndMortyRepository {
	host := "mongo"
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

	return &rickAndMortyRepository{
		client:     client,
		collection: collection,
	}
}

func (r *rickAndMortyRepository) Add(character entities.Character) (id primitive.ObjectID, err error) {
	result, err := r.collection.InsertOne(context.TODO(), character)

	if err != nil {
		log.Fatal(err)
	}

	return result.InsertedID.(primitive.ObjectID), nil
}