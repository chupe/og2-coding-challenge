package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initialize(client *mongo.Client) *mongo.Client {
	dbName := os.Getenv("DB_NAME")
	collection := client.Database(dbName).Collection("Users")

	_, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "code", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatalf("unable to create index: %s", err.Error())
	}

	return client
}
