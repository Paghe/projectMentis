// Package main is the entry point for the Mentis application.
package main

import (
	"fmt"

	"github.com/Paghe/projectMentis/backend/src/api"
	"github.com/Paghe/projectMentis/backend/src/config"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World 2")

	router := gin.Default()

	config.ConnectDB()
	api.SetupRoutes(router)

	if err := router.Run(); err != nil {
		panic(err)
	}
}
