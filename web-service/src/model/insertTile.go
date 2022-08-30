package model

import (
	"strconv"
	"github.com/gomodule/redigo/redis"
	"github.com/leoantony72/rplace/src/config"
	"github.com/leoantony72/rplace/src/utils"
)

func Insert_Tile(offset, value int64) int64 {

	conn := config.NPool()
	var off string = strconv.FormatInt(offset, 10)
	t := "#" + off

	_, err := conn.Do(
		"BITFIELD",
		"place",
		"SET",
		"u8",
		t,
		value,
	)
	utils.CheckErr(err)
	reply, err := redis.Values(conn.Do(
		"BITFIELD",
		"place",
		"GET",
		"u8",
		t,
	))
	utils.CheckErr(err)

	defer conn.Close()

	return reply[0].(int64)

}
