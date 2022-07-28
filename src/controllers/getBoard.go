package controllers

import (
	// "encoding/base64"
	"github.com/leoantony72/rplace/src/services"
	"github.com/gin-gonic/gin"
)

func Get_Board(c *gin.Context) {

	reply := services.Get_Board()
	c.JSON(200, gin.H{"message": "Good duck Processing that 😉", "data": reply})
}



// @Used to convert base64 string into unit8array

	// b, err := base64.StdEncoding.DecodeString("")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%v", b)