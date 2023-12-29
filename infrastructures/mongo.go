package infrastructures

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoDB(url string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalf("failed to connect mongo: %s\n", err.Error())
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("failed to ping mongo: %s\n", err.Error())
	}

	return client
}

func NewMongoDBReadPrefer(url string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url), options.Client().SetReadPreference(readpref.Secondary()))
	if err != nil {
		log.Fatalf("failed to connect mongo: %s\n", err.Error())
	}

	if err := client.Ping(ctx, readpref.Secondary()); err != nil {
		log.Fatalf("failed to ping mongo: %s\n", err.Error())
	}

	return client
}
