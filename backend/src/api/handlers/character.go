// Package handlers contains HTTP handlers for character-related operations.
package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Paghe/projectMentis/backend/src/config"
	"github.com/Paghe/projectMentis/backend/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetCharacter handles retrieving character information.
func GetCharacter(c *gin.Context) {
	c.JSON(200, "Here is Get Character")
}

// EditCharacter handles updating character information.
func EditCharacter(c *gin.Context) {
	c.JSON(200, "Here is Post Character")
}

//CreateCharacter handles POST /character requests and creates a new character in the database.
func CreateCharacter(c *gin.Context) {

	var input struct {
		Nickame	string	`json:"nickname" binding:"required"`
		UserID	string	`json:"user_id" binding:"required"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := primitive.ObjectIDFromHex(input.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	character := models.Character{
		ID:				primitive.NewObjectID(),
		UserID:			userID,
		Nickname:		input.Nickame,
		Level:			1,
		Experience: 	0,
		NextLevelXP:	100,

		BaseStats:	models.NewBaseStats(),

		Gold:		0,
		Gems: 		0,

		TotalTasksCompleted:	0,
		CurrentStreak:			0,
		LongestStreak: 			0,
		LastActiveDate: 		 time.Now(),
		
		CreatedAt:	time.Now(),
		UpdatedAt: 	time.Now(),

	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := config.GetCollection("character")
	_, err = collection.InsertOne(ctx, character)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "character loaded ✅",
		"characters info": character,
	})
}

//GetCharacterByID handles Get /character requests and return character by its ID.
func GetCharacterByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	characterCollection := config.GetCollection("character")
	var character models.Character
	err = characterCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&character)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}
	c.JSON(http.StatusOK, character)
}