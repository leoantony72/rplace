package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/leoantony72/rplace/src/controllers"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("static/index.html")
	router.Static("/css","static/css")
	router.Static("/js","static/js")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.GET("/",controllers.Home)
	router.GET("/ping", controllers.Ping)
	router.GET("/board", controllers.Get_Board)
	router.POST("/tile", controllers.Get_Tile)
	// router.POST("/test", controllers.Test)

	router.Run("0.0.0.0:8080")

}
