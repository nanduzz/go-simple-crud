package repository

import (
	"context"
	"log"

	"github.com/nanduzz/go-simple-crud/entity"
	"github.com/nanduzz/go-simple-crud/infra/db/mongo_repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserRepositoryInstance *UserRepository
var collection *mongo.Collection

// functions for testing

func init() {
	log.Println("init user repository")

	collection = mongo_repository.MongoDatababse.Collection("users")
}

func FindOne() (*entity.User, error) {
	var result entity.User
	err := collection.FindOne(context.TODO(), bson.M{}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func Create(u *entity.User) (*entity.User, error) {
	_, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func FindAll(query map[string]interface{}) ([]*entity.User, error) {
	if query == nil {
		query = bson.M{}
	}

	cursor, err := collection.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}

	var users []*entity.User
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func FindUserById(id string) (*entity.User, error) {
	objejctId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var result entity.User
	err = collection.FindOne(context.TODO(), bson.M{"_id": objejctId}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) Save(u *entity.User) (*entity.User, error) {
	return u, nil
	// return u, errors.New("not implemented")
}
