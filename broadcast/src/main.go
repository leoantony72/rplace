package main

// import "fmt"
import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/leoantony72/rplace/broadcast/src/services"
)

func main() {
	router := gin.Default()

	router.GET("/test", test)
	router.Run("localhost:8081")
}

func test(c *gin.Context) {

	ctx := context.Background()
	services.Consume(c.Writer, c.Request,ctx)
}
