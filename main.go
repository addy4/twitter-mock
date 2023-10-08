package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/gorilla/websocket"
	"projects.com/apps/twitter-app/apis"
	"projects.com/apps/twitter-app/data"
	"projects.com/apps/twitter-app/utils"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	// WebSocket APIs Service
	http.HandleFunc("/web", service)
	http.HandleFunc("/webs", utils.Authorize(service))

	// Broadcast
	go apis.Broadcast()

	// Validate Token
	utils.ValidateToken("abcd")

	// Starting Service
	log.Print("Twitter Server Starting At localhost: 5020")
	if err := http.ListenAndServe(":5020", nil); err != nil {
		log.Fatal(err)
	}
}

func service(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("connection error: " + err.Error())
	}

	data.Clients = append(data.Clients, conn)

	// Adding New Client Session
	session_id := uuid.New().String()
	session := data.ActiveUser{Websocket_Session: conn, Session_Id: data.SessionID(session_id), User_Id: data.ClientID(r.Header.Get("Tw-Client-Id"))}
	data.ActiveSessions[data.SessionID(session_id)] = session
	data.SessionNotifier <- session

	api := &data.RequestDecode{}
	common := &data.CommonAPI{}

	for {

		// Read Message Sent By Client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read error: " + err.Error())
			delete(data.ActiveSessions, data.SessionID(session_id))
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
