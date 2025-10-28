// Package handlers contains HTTP handlers for task-related operations.
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
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateTask handles the creation of a new task.
func CreateTask(c *gin.Context) {
	var input struct {
		UserID		string	`json:"user_id" binding:"required"`
		CharacterID	string	`json:"character_id" binding:"required"`
		Title		string	`json:"title" binding:"required"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UserID, err := primitive.ObjectIDFromHex(input.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID"})
		return
	}

	CharacterID, err := primitive.ObjectIDFromHex(input.CharacterID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CharacterID"})
		return
	}

	task := models.Task {
		ID:				primitive.NewObjectID(),
		UserID: 		UserID,
		CharacterID:	CharacterID,

		Title: input.Title,
		Completed: false,

		CreatedAt: time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := config.GetCollection("task")
	_, err = collection.InsertOne(ctx, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Task registred ✅",
		"data": task,
	})

}

// EditTask handles updating an existing task.
func EditTask(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var input struct {
		Title string `json:"title"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	taskcollection := config.GetCollection("task")
	update := bson.M{
		"$set": bson.M{
			"title": input.Title,
		},
	}
	var updateTask models.Task
	result := taskcollection.FindOneAndUpdate(ctx, 
		bson.M{"_id": objID}, update, 									
		options.FindOneAndUpdate().SetReturnDocument(options.After),
		)
	if err != result.Decode(&updateTask) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"succes": true,
		"message": "Task Updated",
		"data": updateTask,
	})
}

// DeleteTask handles the deletion of a task.
func DeleteTask(c *gin.Context) {
	c.JSON(200, "Here is DeleteTask")
}

// MarkTask handles marking a task as completed.
func MarkTask(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var input struct {
		Completed	bool `json:"completed"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	taskcollection := config.GetCollection("task")
	update := bson.M{
		"$set": bson.M{
			"completed": input.Completed,
		},
	}
	var updateTask models.Task
	result := taskcollection.FindOneAndUpdate(ctx, 
		bson.M{"_id": objID}, update, 									
		options.FindOneAndUpdate().SetReturnDocument(options.After),
		)
	if err != result.Decode(&updateTask) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"succes": true,
		"message": "Task Updated",
		"data": updateTask,
	})
}

// ListTasks handles retrieving a list of tasks.
func ListTasks(c *gin.Context) {
	c.JSON(200, "Here is MarkTask")
}
