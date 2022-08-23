package controllers

import (
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.IndentedJSON(200, gin.H{"message": "Pong!"})
}
