package model

import (
	"github.com/leoantony72/rplace/src/config"
	"github.com/leoantony72/rplace/src/utils"
)

func Flush() {
	conn := config.NPool()

	defer conn.Close()

	_, err := conn.Do(
		"DEL",
		"place",
	)
	utils.CheckErr(err)
}
