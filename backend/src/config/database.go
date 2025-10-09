// Package config provides database configuration for the Mentis application.
package config

import (
	"context"
	"time"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client is the global MongoDB client for the application.
var Client *mongo.Client

// DBname is the name of the MongoDB database used in Mentis.
const DBname = "Mentis"

// ConnectMongo establishes a connection to the MongoDB database.
func ConnectMongo(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	if err = client.Ping(ctx, nil); err != nil {
		return err
	}
	Client = client
	fmt.Println("✅ Connected & pinged MongoDB")
	return nil
}

// GetCollection returns a handle to a MongoDB collection by name.
func GetCollection(name string) *mongo.Collection {
	return Client.Database(DBname).Collection(name)
}
