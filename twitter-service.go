package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Friends
var (
	friends = make(map[int]map[int]bool)
)

// Generic
type APIRequest interface {
	GetAction() string
}

// Follow request
type FollowReq struct {
	Action   string `json:"action"`
	User     int    `json:"userId"`
	Followee int    `json:"followeeId"`
}

// Tweet post
type TweetPost struct {
	Action  string `json:"action"`
	User    int    `json:"userId"`
	Content string `json:"content"`
}

// Implement Interface
func (r FollowReq) GetAction() string {
	return r.Action
}

func (r TweetPost) GetAction() string {
	return r.Action
}

// Clients connected to Twitter Service
var clients = make(map[*websocket.Conn]bool)

// Notifications
var notifier = make(chan FollowReq)

// Upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	fmt.Println("##### Launching Twitter Server....")

	http.HandleFunc("/proj", service)

	go notifications()

	log.Print("Server starting at localhost: 5020")
	if err := http.ListenAndServe(":5020", nil); err != nil {
		log.Fatal(err)
	}
}

func notifications() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		for {
			msg := <-notifier
			fmt.Printf("Request by %d for %d", msg.User, msg.Followee)
			for client := range clients {
				client.WriteJSON(&msg)
			}
		}
	}()

	wg.Done()

}

func follow(friends map[int]map[int]bool, userId int, followee int) {

	if friends[userId] == nil {
		friends[userId] = make(map[int]bool)
	}

	friends[userId][followee] = true
}

func getFollowers(friends map[int]map[int]bool, userID int) {
	fmt.Printf("Getting for %d\n", userID)
	for follower, status := range friends[userID] {
		fmt.Printf("....Follower %d, Status %v\n", follower, status)
	}
}

func service(w http.ResponseWriter, r *http.Request) {

	var wg sync.WaitGroup

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}
	clients[ws] = true

	wg.Add(1)

	go func() {
		for {
			fmt.Println("running2")
			//var request APIRequest
			var request FollowReq
			ws.ReadJSON(&request)
			fmt.Println("test")
			fmt.Println(request)
			//follow(friends, request.GetUser(), request.GetFollowee())
			follow(friends, request.User, request.Followee)
			getFollowers(friends, 1)
			notifier <- request
		}
	}()

	wg.Wait()

}
