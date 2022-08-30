package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username       string             `json:"username" bson:"username" validate:"required,alphanum" example:"example123"`
	IronSpending   int                `json:"iron" bson:"iron" validate:"numeric" example:"42"`
	CopperSpending int                `json:"copper" bson:"copper" validate:"numeric" example:"42"`
	GoldSpending   int                `json:"gold" bson:"gold" validate:"numeric" example:"42"`
	Created        time.Time          `json:"created" bson:"created" validate:"required" example:"2021-05-25T00:00:00.0Z"`
	IronFactory    Factory            `json:"ironFactory" bson:"ironFactory"`
	CopperFactory  Factory            `json:"copperFactory" bson:"copperFactory"`
	GoldFactory    Factory            `json:"goldFactory" bson:"goldFactory"`
}

func (u *User) GetIronOre() int {
	return u.IronFactory.OreProduced() - u.IronSpending
}

func (u *User) GetCopperOre() int {
	return u.CopperFactory.OreProduced() - u.CopperSpending
}

func (u *User) GetGoldOre() int {
	return u.GoldFactory.OreProduced() - u.GoldSpending
}
