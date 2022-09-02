package models

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/chupe/og2-coding-challenge/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	coll *mongo.Collection
}

func NewUsers(env *config.Env) *Users {
	return &Users{
		coll: env.DB.Database(env.Cfg.DB.Name).Collection("Users"),
	}
}

func (u *Users) FindAll() ([]*User, error) {
	var Users []*User
	cursor, err := u.coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, errors.New("error searching collection")
	}

	for cursor.Next(context.TODO()) {
		var User User
		err := cursor.Decode(&User)
		if err != nil {
			log.Fatal(err)
		}
		if err := cursor.Err(); err != nil {
			return nil, err
		}

		Users = append(Users, &User)
	}
	return Users, nil
}

func (u *Users) Find(id string) (*User, error) {
	var User *User
	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = u.coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: bsonId}}).Decode(&User)
	if err != nil {
		return nil, errors.New("error searching collection")
	}

	return User, nil
}

func (u *Users) FindByUsername(username string) (*User, error) {
	var User *User
	err := u.coll.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&User)
	if err != nil {
		return nil, err
	}

	return User, nil
}

func (u *Users) Create(username string) (*User, error) {
	User := &User{
		Username:      username,
		IronFactory:   NewIronFactory(),
		CopperFactory: NewCopperFactory(),
		GoldFactory:   NewGoldFactory(),
		Created:       time.Now().UTC(),
	}

	result, err := u.coll.InsertOne(context.TODO(), User)
	if err != nil {
		return User, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		User.ID = oid
	}

	return User, nil
}

func (u *Users) Update(user *User) error {
	_, err := u.coll.ReplaceOne(context.TODO(), bson.D{{"_id", user.ID}}, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *Users) Delete(id string) (string, error) {
	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	_, err = u.coll.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: bsonId}})
	if err != nil {
		return bsonId.String(), err
	}

	return bsonId.String(), nil
}
