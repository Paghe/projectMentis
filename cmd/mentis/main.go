// Package main is the entry point for the Mentis application.
package main

import (
	"fmt"
	"log"                    // For logging errors
	"github.com/Paghe/projectMentis/backend/src/api"
	"github.com/Paghe/projectMentis/backend/src/config"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.ConnectMongo("mongodb://localhost:27017"); err != nil {
		log.Fatal("❌ MongoDB connection failed: ", err)
	}
	fmt.Println("✅ Database ready!")

	router := gin.Default()
	api.SetupRoutes(router)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
