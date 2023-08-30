package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"projects.com/apps/twitter-app/apis"
	"projects.com/apps/twitter-app/data"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	// WebSocket APIs Service
	http.HandleFunc("/web", service)

	// Broadcast
	go apis.Broadcast()

	// Starting Service
	log.Print("Twitter Server Starting At localhost: 5020")
	if err := http.ListenAndServe(":5020", nil); err != nil {
		log.Fatal(err)
	}
}

func service(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}

	data.Clients = append(data.Clients, conn)

	api := &data.RequestDecode{}
	common := &data.CommonAPI{}

	for {

		// Read Message Sent By Client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read error: " + err.Error())
			conn.Close()
			return
		}

		// Unmarshall JSON For Decoding
		err = json.Unmarshal(msg, &api)
		if err != nil {
			fmt.Println("unmarshal error: " + err.Error())
		}

		// Unmarshall JSON For Getting Request Type
		json.Unmarshal(msg, &common)

		// Handle Request
		handler := apis.ActionHandlers[common.Action]
		handler(conn, api)
	}
}
