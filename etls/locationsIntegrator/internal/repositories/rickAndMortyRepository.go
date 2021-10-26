package Repositories

type RickAndMortyRepository interface {
	RepositoryBase
}

type rickAndMortyRepository struct {
	repositoryBase
}

func NewRickAndMortyRepository() *rickAndMortyRepository {
	return &rickAndMortyRepository{
		repositoryBase: *NewRepositoryBase("Locations"),
	}
}

// package Repositorys

// import (
// 	"context"
// 	"fmt"
// 	entities "locationsIntegrator/internal"
// 	"log"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	mongo "go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type rickAndMortyRepository struct {
// 	client     *mongo.Client
// 	collection *mongo.Collection
// }

// func NewRickAndMortyRepository() *rickAndMortyRepository {
// 	host := "mongo"
// 	port := 27017

// 	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
// 	client, err := mongo.Connect(context.TODO(), clientOpts)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Check the connections
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Congratulations, you're already connected to MongoDB!")
// 	collection := client.Database("RickAndMorty").Collection("Locations")

// 	return &rickAndMortyRepository{
// 		client:     client,
// 		collection: collection,
// 	}
// }

// func (r *rickAndMortyRepository) Add(entity entities.Location) (id primitive.ObjectID, err error) {
// 	result, err := r.collection.InsertOne(context.TODO(), entity)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return result.InsertedID.(primitive.ObjectID), nil
// }
