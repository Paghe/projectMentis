package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the Mentis application.
type User struct {
	ID			primitive.ObjectID	`bson:"_id,omitempty" json:"id"`
	Email		string				`bson:"email" json:"email"`
	Username	string				`bson:"username" json:"username"`
	CharacterID primitive.ObjectID	`bson:"character_id,omitempty" json:"character_id"`
	CreatedAt	time.Time			`bson:"created_at" json:"created_at"`
	UpdatedAt	time.Time			`bson:"updated_at" json:"updated_at"`
}
