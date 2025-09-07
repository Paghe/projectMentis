package api

import (
	"github.com/gin-gonic/gin"
)

func GetTask(c* gin.Context) {
	c.JSON(200, "Here is Get")
}

func PostTask(c* gin.Context) {
	c.JSON(200, "Here is Post")
}

func PutTask(c* gin.Context) {
	c.JSON(200, "Here is Put")
}
