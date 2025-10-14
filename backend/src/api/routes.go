// Package api provides HTTP routing and handlers for the Mentis application.
package api

import (
	"github.com/Paghe/projectMentis/backend/src/api/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the application.
func SetupRoutes(router *gin.Engine) {
	router.GET("/task", handlers.CreateTask)
	router.POST("/task", handlers.EditTask)
	router.PUT("/task", handlers.DeleteTask)
	router.POST("/users", handlers.CreateUserHandle)
	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUserByID)
	router.POST("/baseStats", handlers.CreateBaseStatsHandle)
	router.POST("/character", handlers.CreateCharacter)
}
