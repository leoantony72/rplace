package model

import (
	"github.com/leoantony72/rplace/src/config"
	"github.com/leoantony72/rplace/src/utils"
)

func Get_Board() interface{} {
	conn := config.NPool()
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
	utils.CheckErr(err)
	return reply
}

