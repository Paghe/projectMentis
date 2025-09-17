package main

import (
	"fmt"
	"github.com/Paghe/projectMentis/backend/src/api"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World 2")

	router := gin.Default()

	api.SetupRoutes(router)
	
	router.Run()
}
