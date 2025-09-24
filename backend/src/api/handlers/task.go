// Package handlers contains HTTP handlers for task-related operations.
package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/Paghe/projectMentis/backend/src/config"
	"github.com/Paghe/projectMentis/backend/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ListTasks handles retrieving a list of tasks.
func ListTasks(c *gin.Context) {
	userId := c.Param("userId")

	collection := config.DB.Collection("tasks")

	filter := bson.M{"userId": userId}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	defer func() {
		if err := cursor.Close(context.TODO()); err != nil {
			log.Printf("Warning: failed to close cursor: %v", err)
		}
	}()

	var tasks []models.Task

	if err := cursor.All(context.TODO(), &tasks); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, tasks)
}

// CreateTask handles the creation of a new task.
func CreateTask(c *gin.Context) {
	var task models.Task

	fmt.Println("I got called")
	if err := c.ShouldBindJSON(&task); err != nil {
		fmt.Println("I got Error")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	collection := config.DB.Collection("tasks")
	result, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(201, gin.H{"id": result.InsertedID})
}

// EditTask handles updating an existing task.
func EditTask(c *gin.Context) {
	userId := c.Param("userId")
	taskId := c.Param("taskId")

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	collection := config.DB.Collection("tasks")

	objectId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	filter := bson.M{
		"userId": userId,
		"_id":    objectId,
	}

	update := bson.M{"$set": updateData}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(404, gin.H{"error": "Task not found or no changes made"})
		return
	}

	c.JSON(200, gin.H{"message": "Task updated successfully"})
}

// DeleteTask handles the deletion of a task.
func DeleteTask(c *gin.Context) {
	userId := c.Param("userId")
	taskId := c.Param("taskId")

	collection := config.DB.Collection("tasks")
	objectId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	filter := bson.M{
		"userId": userId,
		"_id":    objectId,
	}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(404, gin.H{"error": "Task not found or doesn't belong to user"})
		return
	}

	c.JSON(200, gin.H{"message": "Task deleted successfully", "taskId": taskId})
}
