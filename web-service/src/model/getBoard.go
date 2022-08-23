package model

import (
	// "fmt"

	// "fmt"

	"github.com/leoantony72/rplace/src/config"
	"github.com/leoantony72/rplace/src/utils"
)

func Get_Board() interface{} {
	conn := config.NPool()
	defer conn.Close()
	// reply, err := conn.Do(
	// 	"BITFIELD",
	// 	"place",
	// 	"Get",
	// 	"u8",
	// 	"0",
	// )
	// utils.CheckErr(err)

	reply, err := conn.Do(
		"GET",
		"place",
	)
	// g,err:= redis.
	utils.CheckErr(err)
	// fmt.Println(reply)
	return reply
}
