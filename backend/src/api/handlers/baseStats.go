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

//CreateBaseStatsHandle for the user
func CreateBaseStatsHandle(c *gin.Context) {

	var stats models.BaseStats

	if err := c.BindJSON(&stats); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stats.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := config.GetCollection("baseStats")
	_, err := collection.InsertOne(ctx, stats)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "base stats updated",
		"stats": stats,
	})
}

//GetBaseStatsByID for the stats 
func GetBaseStatsByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	statsCollection := config.GetCollection("baseStats")
	var stats models.BaseStats
	err = statsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&stats)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "stats not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Stats found",
		"baseStats": stats,	
	})
}