// Package models contains data structures for the Mentis application.
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task represents a task in the system.
type Task struct {
	ID			primitive.ObjectID 	`bson:"_id,omitempty" json:"id"`
	UserID		primitive.ObjectID	`bson:"user_id" json:"user_id"`
	CharacterID	primitive.ObjectID	`bson:"character_id" json:"character_id"`
	Title     	string				`bson:"title" json:"title"`
	Completed 	bool				`bson:"completed" json:"completed"`
	CreatedAt   time.Time			`bson:"created_at" json:"created_at"`
}
