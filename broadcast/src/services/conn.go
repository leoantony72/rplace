package services

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/segmentio/kafka-go"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *kafka.Message)


/*func kafka_reader(reader *kafka.Reader, ctx context.Context) {
	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		panic("could not read message " + err.Error())
	}
	fmt.Printf("kafka:  %v\n", msg)
	broadcast <- &msg
}*/

func Wshandler(w http.ResponseWriter, r *http.Request, ctx context.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "my-topic",
		GroupID: "my-group",
	})

	// register client
	clients[conn] = true
	for {

		// the `ReadMessage` method blocks until we receive the next event
		/*****************/
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		broadcast <- &msg
		/*****************/
	}
}

func Echo() {
	for {
		msg := <-broadcast
		// send to every client that is currently connected
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg.Value))
			if err != nil {
				log.Printf("Websocket error: %s", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
