package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/rplace/src/services"
)

func Flush(c *gin.Context) {
	services.Flush()
	c.JSON(200, gin.H{"message": "canvas cleared"})
}
