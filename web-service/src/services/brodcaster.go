package services

import (
	"encoding/json"
	"fmt"

	"github.com/leoantony72/rplace/src/config"
	"github.com/leoantony72/rplace/src/utils"
	"github.com/segmentio/kafka-go"
)

func Brodcast(X, Y, Color int64) {
	type data struct {
		X     int64 `json:"x"`
		Y     int64 `json:"y"`
		Color int64 `json:"color"`
	}
	var client_data = data{X: X, Y: Y, Color: Color}

	//parse the data as json
	jsonstring, err := json.Marshal(client_data)
	utils.CheckErr(err)

	//convert the json to string 
	//for passing into the kafka-go writeMessage
	client_data_string := string(jsonstring)
	fmt.Println(client_data_string)

	conn := config.KafkaConfig()
	d, err := conn.WriteMessages(
		kafka.Message{Value: []byte(client_data_string)},
	)
	utils.CheckErr(err)
	conn.Close()

	fmt.Println(d)
}
