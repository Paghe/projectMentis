// Package api provides HTTP routing and handlers for the Mentis application.
package api

import (
	"github.com/Paghe/projectMentis/backend/src/api/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the application.
func SetupRoutes(router *gin.Engine) {
	router.GET("/task/:userId", handlers.ListTasks)
	router.POST("/task", handlers.CreateTask)
	router.PUT("/task/:userId/:taskId", handlers.EditTask)
	router.DELETE("/task/:userId/:taskId", handlers.DeleteTask)
	// router.POST("/task", handlers.EditTask)
}
