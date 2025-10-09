package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the Mentis application.
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email     string             `bson:"email" json:"email"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
