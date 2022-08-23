package services

import (
	"github.com/leoantony72/rplace/src/model"
)

func Insert_Tile(x, y, color int64) int64 {
	//@formula to convert cords to offest
	var offset int64 = y*1000 + x
	var value int64 = color

	//insert into redis db and returns the color
	color_val := model.Insert_Tile(offset, value)
	return color_val
}
