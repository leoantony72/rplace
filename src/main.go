package main

import (
	"example/web-service-gin/src/controllers"

	"github.com/gin-gonic/gin"

)




func main() {
	router := gin.Default()
	router.GET("/",controllers.Get_Board)

	router.Run("localhost:8080")

}
