package data

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/chupe/og2-coding-challenge/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	coll *mongo.Collection
}

func NewUserRepository(dbClient *mongo.Client) *UserRepository {
	dbName := os.Getenv("DB_NAME")
	return &UserRepository{
		coll: dbClient.Database(dbName).Collection("Users"),
	}
}

func (r *UserRepository) FindAll() ([]*models.User, error) {
	var Users []*models.User
	cursor, err := r.coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, errors.New("error searching collection")
	}

	for cursor.Next(context.TODO()) {
		var User models.User
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

func (r *UserRepository) Find(id string) (*models.User, error) {
	var User *models.User
	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = r.coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: bsonId}}).Decode(&User)
	if err != nil {
		return nil, errors.New("error searching collection")
	}

	return User, nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var User *models.User
	err := r.coll.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&User)
	if err != nil {
		return nil, err
	}

	return User, nil
}

func (r *UserRepository) Create(username string) (*models.User, error) {
	User := &models.User{
		Username:      username,
		IronFactory:   models.NewIronFactory(),
		CopperFactory: models.NewCopperFactory(),
		GoldFactory:   models.NewGoldFactory(),
		Created:       time.Now().UTC(),
	}

	result, err := r.coll.InsertOne(context.TODO(), User)
	if err != nil {
		return User, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		User.ID = oid
	}

	return User, nil
}

func (r *UserRepository) Update(user *models.User) error {
	_, err := r.coll.ReplaceOne(context.TODO(), bson.D{{"_id", user.ID}}, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id string) (string, error) {
	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	_, err = r.coll.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: bsonId}})
	if err != nil {
		return bsonId.String(), err
	}

	return bsonId.String(), nil
}
