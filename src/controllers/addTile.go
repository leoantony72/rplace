package controllers

import (
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/leoantony72/rplace/src/services"
)

func Get_Tile(c *gin.Context) {

	type Body struct {
		X     int64 `form:"x"`
		Y     int64 `form:"y"`
		Color int64 `form:"color"`
	}
	body := Body{}

	//@binds the data from query params to Body struct
	if err := c.ShouldBindQuery(&body); err != nil {
		panic(err)
	}

	//@passes the cordinates and returns color
	color := services.Insert_Tile(body.X, body.Y, body.Color)
	c.JSON(201, gin.H{"status": "success", "color": color})
}
