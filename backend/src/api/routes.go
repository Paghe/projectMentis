// Package api provides HTTP routing and handlers for the Mentis application.
package api

import (
	"github.com/Paghe/projectMentis/backend/src/api/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the application.
func SetupRoutes(router *gin.Engine) {
	router.POST("/task", handlers.CreateTask)
	router.PUT("/task/:id", handlers.EditTask)
	router.PUT("/taskmark/:id", handlers.MarkTask)
	router.DELETE("/task", handlers.DeleteTask)
	router.POST("/users", handlers.CreateUserHandle)
	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUserByID)
	router.POST("/baseStats", handlers.CreateBaseStatsHandle)
	router.POST("/character", handlers.CreateCharacter)
	router.GET("/character/:id", handlers.GetCharacterByID)
}
