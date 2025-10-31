package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Paghe/projectMentis/backend/src/config"
	"github.com/Paghe/projectMentis/backend/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

//GetUserByID handles Get /user requests and return user by its ID.
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	userCollection := config.GetCollection("user")
	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

//UpdateUser update email of the user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Email	string `json:"email"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}
	UserID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	userCollection := config.GetCollection("user")
	update := bson.M{
		"$set": bson.M{
			"email": input.Email,
		},
	}
	var updateUser models.User
	result := userCollection.FindOneAndUpdate(ctx, 
		bson.M{"_id": UserID}, update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
		)
	if err != result.Decode(&updateUser) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User updated",
		"data": updateUser,
	})
}