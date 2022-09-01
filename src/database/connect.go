// Package database deals with creating Mongo client
package database

import (
	"context"
	"fmt"
	"log"

	"github.com/chupe/og2-coding-challenge/config"
	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(cfg *config.DB) *mongo.Client {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database", err.Error())
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database", err.Error())
	}
	color.Green("⛁ Connected to Database")

	initialize(client, cfg)

	return client
}
