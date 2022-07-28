package controllers

import (
	// "encoding/base64"
	"log"

	"example/web-service-gin/src/config"

	"github.com/gin-gonic/gin"
)


func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
//redis function to retrive the whole board
func Get_Board(c *gin.Context) {

	conn := config.NPool()
	reply, err := conn.Do(
		"BITFIELD",
		"place",
		"Get",
		"u8",
		"0",
	)
	checkErr(err)

	replys, err := conn.Do(
		"GET",
		"place",
	)
	checkErr(err)


	c.JSON(200, gin.H{"message": reply, "data": replys})
}
	// b, err := base64.StdEncoding.DecodeString("")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%v", b)