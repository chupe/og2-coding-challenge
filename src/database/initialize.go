package database

import (
	"context"
	"log"

	"github.com/chupe/og2-coding-challenge/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initialize(client *mongo.Client, cfg *config.DB) *mongo.Client {
	collection := client.Database(cfg.Name).Collection("Users")

	_, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatalf("unable to create index: %s", err.Error())
	}

	return client
}
