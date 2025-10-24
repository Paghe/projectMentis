// Package handlers contains HTTP handlers for task-related operations.
package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Paghe/projectMentis/backend/src/config"
	"github.com/Paghe/projectMentis/backend/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	c.JSON(200, "Here is EditTaskTask")
}

// DeleteTask handles the deletion of a task.
func DeleteTask(c *gin.Context) {
	c.JSON(200, "Here is DeleteTask")
}

// MarkTask handles marking a task as completed.
func MarkTask(c *gin.Context) {
	c.JSON(200, "Here is MarkTask")
}

// ListTasks handles retrieving a list of tasks.
func ListTasks(c *gin.Context) {
	c.JSON(200, "Here is MarkTask")
}
