package config

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Env struct {
	DB  *mongo.Client
	Cfg *Config
}
