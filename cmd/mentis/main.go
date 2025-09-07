package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/Paghe/projectMentis/backend/src/api"

)

func main() {
	fmt.Println("Hello World 2")

	router := gin.Default()

	router.GET("/task", api.GetTask)
	router.POST("/task", api.PostTask)
	router.PUT("/task", api.PutTask) 

	router.Run()
}
