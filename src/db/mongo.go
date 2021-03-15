package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var err error

type database struct {
	DB *mongo.Client
}

// GetInstance return copy of db session
func GetInstance() *mongo.Client {

	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		os.Getenv("dbuser"),
		os.Getenv("dbpassword"),
		os.Getenv("dbserver"),
		os.Getenv("dbname"))

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	return client
}
