package model

import (
	"github.com/leoantony72/rplace/src/config"
	"github.com/leoantony72/rplace/src/utils"
)

func Get_Board() interface{} {
	conn := config.NPool()

	defer conn.Close()

	reply, err := conn.Do(
		"GET",
		"place",
	)
	utils.CheckErr(err)
	return reply
}
