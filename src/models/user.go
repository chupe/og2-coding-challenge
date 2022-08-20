package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username" validate:"required,alphanum" example:"example123"`
	Iron     int                `json:"iron" bson:"iron" validate:"numeric" example:"42"`
	Copper   int                `json:"copper" bson:"copper" validate:"numeric" example:"42"`
	Gold     int                `json:"gold" bson:"gold" validate:"numeric" example:"42"`
	Created  time.Time          `json:"created" validate:"required" example:"2021-05-25T00:00:00.0Z"`
}
