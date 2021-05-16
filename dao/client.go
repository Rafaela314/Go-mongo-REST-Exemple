package dao

import (
	"context"
	"os"
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	database = "starwars"
)

// Connect with running mongodb server
func Connect() (*mongo.Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	return client.Database(database, nil), err
}

// ToDocument will convert any struct into bson.D primitive struct
func ToDocument(doc interface{}) *bson.D {
	data, err := bson.Marshal(doc)
	var docD *bson.D
	if err != nil {
		log.Info("toDocument Marshal %+v:", doc)
		return docD
	}
	err = bson.Unmarshal(data, &docD)
	if err != nil {
		log.Info("toDocument Unmarshal %+v", doc)
	}
	return docD
}
