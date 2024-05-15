package mongo_repository

import (
	"context"
	"log"
	"os"

	"github.com/nanduzz/go-simple-crud/util"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDatababse *mongo.Database

// functions for testing
var loadEnvFunc = util.LoadEnv

func init() {
	loadEnvFunc()
	initMongoClient()

}

func initMongoClient() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI is not set")
	}
	credentials := options.Credential{
		AuthSource: os.Getenv("MONGODB_AUTH_SOURCE"),
		Username:   os.Getenv("MONGODB_USER"),
		Password:   os.Getenv("MONGODB_PASS"),
	}

	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			log.Print(evt.Command)
		},
	}
	clientOptionso := options.Client().ApplyURI(uri).SetAuth(credentials).SetMonitor(cmdMonitor)
	mongoClient, err := mongo.Connect(context.TODO(), clientOptionso)
	if err != nil {
		log.Fatal(err)
	}

	database := os.Getenv("MONGODB_DATABASE")
	MongoDatababse = mongoClient.Database(database)
}
