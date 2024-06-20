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
	router.Run("0.0.0.0:8081")
}

func test(c *gin.Context) {

	go services.Echo()
	ctx := context.Background()
	services.Wshandler(c.Writer, c.Request, ctx)

}
