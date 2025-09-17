// Package config provides database configuration for the Mentis application.
package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB is the global database connection for the Mentis application.
var DB *mongo.Database

// ConnectDB establishes a connection to the MongoDB database.
func ConnectDB() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	DB = client.Database("mentis")
}
