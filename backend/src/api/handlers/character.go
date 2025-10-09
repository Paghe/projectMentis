// Package handlers contains HTTP handlers for character-related operations.
package handlers

import (
	"github.com/gin-gonic/gin"
)

// GetCharacter handles retrieving character information.
func GetCharacter(c *gin.Context) {
	c.JSON(200, "Here is Get Character")
}

// EditCharacter handles updating character information.
func EditCharacter(c *gin.Context) {
	c.JSON(200, "Here is Post Character")
}
