package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"projects.com/apps/twitter-app/apis"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	msg := &apis.RequestDecode{}
	jsonData := `
		{
			"follow" : {
				"current_user": 1,
				"followed_user": 2
			}
		}
	`
	err := json.Unmarshal([]byte(jsonData), &msg)

	fmt.Println(err)
	fmt.Println(msg.FollowRequestDetails)
	fmt.Print(msg.PostRequestDetails)

	fmt.Println("##### Launching Twitter Server....")

	http.HandleFunc("/proj", service)

	log.Print("Server starting at localhost: 5020")
	if err := http.ListenAndServe(":5020", nil); err != nil {
		log.Fatal(err)
	}
}

func service(w http.ResponseWriter, r *http.Request) {

	var wg sync.WaitGroup

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	api := &apis.RequestDecode{}
	common := &apis.CommonAPI{}

	wg.Add(1)

	go func() {
		for {
			_, msg, err := conn.ReadMessage()

			if err != nil {
				fmt.Println("Error: " + err.Error())
			}

			err = json.Unmarshal(msg, &api)

			if err != nil {
				fmt.Println("Error: " + err.Error())
			}

			err = json.Unmarshal(msg, &common)

			handler := apis.ActionHandlers[common.Action]
			fmt.Println("here")
			fmt.Println(handler)
			a := handler(conn, api)
			fmt.Println(a)
			//apis.FollowRequest(conn, api)
		}
	}()

	wg.Wait()
}
