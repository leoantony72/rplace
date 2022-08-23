package services

import "fmt"
import "net/http"
import "github.com/gorilla/websocket"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Wshandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	
	type data struct {
		Name string `json:name`
		Age  int64  `json:age`
	}

	var st data
	st.Name="leo"
	st.Age=17

	// msg := []byte("string")
	err = conn.WriteJSON(st)
	fmt.Println(err)
}
