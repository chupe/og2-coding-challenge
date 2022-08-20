// Package database deals with creating Mongo client
package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DbClient return a *mongo.Client
func DbClient() *mongo.Client {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	mongoUrl := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", user, password, host, port)
	clientOptions := options.Client().ApplyURI(mongoUrl)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database", err.Error())
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database", err.Error())
	}
	color.Green("⛁ Connected to Database")

	initialize(client)

	return client
}
