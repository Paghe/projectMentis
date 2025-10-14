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
	var character models.Character

	if err := c.BindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if character.Nickname == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nickname is required"})
		return
	}
	character.ID = primitive.NewObjectID()
	character.Level = 1
	character.Experience = 0
	character.CreatedAt = time.Now()
	character.UpdatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := config.GetCollection("character")
	_, err := collection.InsertOne(ctx, character)
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