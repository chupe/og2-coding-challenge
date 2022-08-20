package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chupe/og2-coding-challenge/domain"
	"github.com/chupe/og2-coding-challenge/models"
	"github.com/go-playground/validator/v10"
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
	bsonId, _ := primitive.ObjectIDFromHex(id)
	err := r.coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: bsonId}}).Decode(&User)
	if err != nil {
		return nil, errors.New("error searching collection")
	}

	return User, nil
}

func (r *UserRepository) FindByCode(code string) (*models.User, error) {
	var User *models.User
	err := r.coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "code", Value: code}}).Decode(&User)
	if err != nil {
		return nil, err
	}
	err = r.incrementHitCount(User.ID)
	if err != nil {
		return nil, err
	}

	return User, nil
}

func (r *UserRepository) Create(url string) (*models.User, error) {
	User := &models.User{
		Url:      url,
		HitCount: 0,
		Code:     domain.GenerateCode(6),
		Created:  time.Now().UTC(),
	}

	v := validator.New()
	err := v.Struct(User)
	if err != nil {
		msg := "Validation failed"
		for _, e := range err.(validator.ValidationErrors) {
			msg = fmt.Sprintf("%s, %s", msg, e)
		}
		return nil, errors.New(msg)
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

func (r *UserRepository) incrementHitCount(UserId primitive.ObjectID) error {
	res, err := r.coll.UpdateOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: UserId}}, bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "hitCount", Value: 1}}}})
	if err != nil {
		return err
	}
	if res.MatchedCount != 1 {
		return errors.New("failed to increment hit count")
	}

	return nil
}

func (r *UserRepository) Delete(UserId primitive.ObjectID) (string, error) {
	_, err := r.coll.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: UserId}})
	if err != nil {
		return UserId.String(), err
	}

	return UserId.String(), nil
}
