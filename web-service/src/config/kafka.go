package config

import (
	"context"

	"github.com/leoantony72/rplace/src/utils"
	"github.com/segmentio/kafka-go"
)

func KafkaConfig() *kafka.Conn {
	topic := "my-topic"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "127.0.0.1:9092", topic, partition)

	utils.CheckErr(err)

	return conn
}
