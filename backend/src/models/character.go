package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Character represents a character in the system.
type Character struct {
	ID			primitive.ObjectID	`bson:"_id,omitempty" json:"id"`
	UserID 		primitive.ObjectID	`bson:"user_id" json:"user_id"`
	Nickname	string				`bson:"nickname" json:"nickname"`
	Level		int					`bson:"level" json:"level"`
	NextLevelXP int             	`bson:"next_level_xp" json:"next_level_xp"`
	Experience	int					`bson:"experience" json:"experience"`
	CreatedAt	time.Time			`bson:"created_at" json:"created_at"`
	UpdatedAt	time.Time			`bson:"updated_at" json:"updated_at"`
	TotalTasksCompleted	int			`bson:"total_tasks_completed" json:"total_tasks_completed"`
	CurrentStreak		int       	`bson:"current_streak" json:"current_streak"`
	LongestStreak		int       	`bson:"longest_streak" json:"longest_streak"`
	LastActiveDate      time.Time	`bson:"last_active_date" json:"last_active_date"`
	BaseStats	BaseStats			`bson:"base_stats" json:"base_stats"`
	Gold		int					`bson:"gold" json:"gold"`
	Gems		int					`bson:"gems" json:"gems"`
}
