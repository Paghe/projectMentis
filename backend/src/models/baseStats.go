// Package models contains data structures for the Mentis application.
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BaseStats rapperenst base state of the user
type BaseStats struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Strength     int                `json:"strength" bson:"strength"`
	Stamina      int                `json:"stamina" bson:"stamina"`
	Dexterity    int                `json:"dexterity" bson:"dexterity"`
	Intelligence int                `json:"intelligence" bson:"intelligence"`
	Constitution int                `json:"constitution" bson:"constitution"`
}

//NewBaseStats cunstructor for BaseStats
func NewBaseStats() BaseStats {
	return BaseStats{
		Strength: 		10,           
		Stamina:		10,
		Dexterity:		10,
		Intelligence:	10,
		Constitution:	10,
	}
}
