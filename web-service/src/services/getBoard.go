package services

import "github.com/leoantony72/rplace/src/model"

func Get_Board() interface{} {

	reply := model.Get_Board()
	return reply
}
