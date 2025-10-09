package handlers

import (
	"context"
	"net/http"
	"time"
	"fmt"
	"github.com/Paghe/projectMentis/backend/src/config"
	"github.com/Paghe/projectMentis/backend/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreateUserHandle handles POST /users requests and creates a new user in the database.
func CreateUserHandle(c *gin.Context) {
	
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := config.GetCollection("user")
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
		"user": user,
	})
}

//GetUsers handles GET /users requests and returns all users in the database.
func GetUsers(c *gin.Context) {
	var users []models.User
	userCollection := config.GetCollection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	results, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func() {
		if err := results.Close(ctx); err != nil {
			fmt.Println("❌ Error closing cursor:", err)
		}
	}()
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		users = append(users, singleUser)
	}
	fmt.Println("✅ Total users found:", len(users))
	c.JSON(http.StatusOK, gin.H{"users": users})
}
