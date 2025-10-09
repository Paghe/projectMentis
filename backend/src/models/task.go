// Package models contains data structures for the Mentis application.
package models

// Task represents a task in the system.
type Task struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Title     string `json:"title" bson:"title"`
	Completed bool   `json:"completed" bson:"completed"`
}
