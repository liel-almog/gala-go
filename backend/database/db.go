package database

import (
	"context"
	"sync"

	"github.com/liel-almog/gala-go/backend/configs"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoClient struct {
	Client    *mongo.Client
	EventColl *mongo.Collection
	GuestColl *mongo.Collection
}

var (
	db         *MongoClient
	initDBOnce sync.Once
)

func newDB() {
	initDBOnce.Do(func() {
		uri, err := configs.GetEnv("DATABASE_URL")
		if err != nil {
			panic(err)
		}

		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

		client, err := mongo.Connect(opts)
		if err != nil {
			panic(err)
		}

		dbName, err := configs.GetEnv("DB_NAME")
		if err != nil {
			panic(err)
		}

		eventColl := client.Database(dbName).Collection("events")
		guestColl := client.Database(dbName).Collection("guests")

		db = &MongoClient{
			Client:    client,
			EventColl: eventColl,
			GuestColl: guestColl,
		}
	})
}

func (c *MongoClient) Close(ctx context.Context) error {
	return c.Client.Disconnect(ctx)
}

func GetDB() *MongoClient {
	if db == nil {
		newDB()
	}

	return db
}
