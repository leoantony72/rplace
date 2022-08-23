package main

// import "fmt"
import "github.com/gin-gonic/gin"
import "github.com/leoantony72/rplace/broadcast/src/services"

func main() {
	router := gin.Default()

	router.GET("/test", test)
	router.Run("localhost:8081")
}

func test(c *gin.Context) {

	services.Wshandler(c.Writer, c.Request)
}
