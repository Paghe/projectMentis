// Package models contains data structures for the Mentis application.
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Task represents a task in the system.
type Task struct {
	ID			string 				`bson:"_id,omitempty" json:"id"`
	UserID		primitive.ObjectID	`bson:"user_id" json:"user_id"`
	Character	primitive.ObjectID	`bson:"character_id" json:"character_id"`
	Title     	string `json:"title" bson:"title"`
	Completed 	bool   `json:"completed" bson:"completed"`
}
