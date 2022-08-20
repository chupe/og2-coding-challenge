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
	bsonId, _ := primitive.ObjectIDFromHex(id)
	err := r.coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: bsonId}}).Decode(&User)
	if err != nil {
		return nil, errors.New("error searching collection")
	}

	return User, nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var User *models.User
	err := r.coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "username", Value: username}}).Decode(&User)
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

func (r *UserRepository) UpgradeFactory(username, factory string) (*models.User, error) {
	user, err := r.FindByUsername(username)
	if err != nil {
		return user, err
	}

	err = deduceOres(user, factory)
	if err != nil {
		return nil, err
	}
	upgradeFactory(user, factory)

	_, err = r.coll.ReplaceOne(context.TODO(), bson.D{{"username", username}}, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func deduceOres(user *models.User, factory string) error {
	cost := models.Ores{}
	switch factory {
	case "iron":
		facLvl := user.IronFactory.GetLevel()
		cost = models.IronConfig.Info[facLvl-1].Cost
	case "copper":
		facLvl := user.CopperFactory.GetLevel()
		cost = models.CopperConfig.Info[facLvl-1].Cost
	case "gold":
		facLvl := user.CopperFactory.GetLevel()
		cost = models.CopperConfig.Info[facLvl-1].Cost
	}

	user.IronSpending += cost.Iron
	user.CopperSpending += cost.Copper
	user.GoldSpending += cost.Gold

	if user.GetIronOre() < 0 || user.GetCopperOre() < 0 || user.GetGoldOre() < 0 {
		return errors.New("not enough resources")
	}

	return nil
}

func upgradeFactory(user *models.User, factory string) {
	switch factory {
	case "iron":
		fac := &user.IronFactory
		lvl := fac.GetLevel()
		fac.UpgradeData[lvl] = time.Now().UTC().Add(time.Second * time.Duration(models.IronConfig.Info[lvl-1].UpgradeDuration))
	case "copper":
		fac := &user.IronFactory
		lvl := fac.GetLevel()
		fac.UpgradeData[lvl] = time.Now().UTC().Add(time.Second * time.Duration(models.IronConfig.Info[lvl-1].UpgradeDuration))
	case "gold":
		fac := &user.IronFactory
		lvl := fac.GetLevel()
		fac.UpgradeData[lvl] = time.Now().UTC().Add(time.Second * time.Duration(models.IronConfig.Info[lvl-1].UpgradeDuration))
	}
}

func (r *UserRepository) Delete(UserId primitive.ObjectID) (string, error) {
	_, err := r.coll.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: UserId}})
	if err != nil {
		return UserId.String(), err
	}

	return UserId.String(), nil
}
