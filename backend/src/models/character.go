package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Character represents a character in the system.
type Character struct {
	ID primitive.ObjectID	`bson:"_id,omitempty" json:"id"`
	OwnerRef string			`bson:"owner_ref" json:"owner_ref"`
	Nickname string			`bson:"nickname" json:"nickname"`
	Level	int				`bson:"level" json:"level"`
	Experience int			`bson:"experience" json:"experience"`
	CreatedAt time.Time		`bson:"created_at" json:"created_at"`
	UpdatedAt time.Time		`bson:"updated_at" json:"updated_at"`
}
