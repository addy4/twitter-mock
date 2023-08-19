package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Generic
type APIRequest interface {
	GetAction() string
	GetUser() int
	GetContent() string
	GetFollowee() int
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

func (r FollowReq) GetUser() int {
	fmt.Println("get user")
	fmt.Println(r.User)
	return r.User
}

func (r FollowReq) GetFollowee() int {
	return r.Followee
}

func (r TweetPost) GetUser() int {
	return r.User
}

func (r TweetPost) GetContent() string {
	return r.Content
}

// Clients connected to Twitter Service
var clients = make(map[*websocket.Conn]bool)

// Notifications
var notifier = make(chan FollowReq)

//var notifier = make(chan APIRequest)
var posts = make(chan TweetPost)
var notes = make(chan FollowReq)

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

	go func() {
		for {
			msg := <-notifier
			fmt.Printf("Request by %d for %d", msg.GetUser(), msg.GetFollowee())
			for client := range clients {
				client.WriteJSON(&msg)
			}
		}
	}()

	for {
		//post := <-posts
		//note := <-notifier
	}
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
	friends := make(map[int]map[int]bool) // every connection request leads to a new memory space, hence inconsistency!,make it global
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	clients[ws] = true

	/*
		go func() {
			for {
				fmt.Println("running2")
				var request FollowReq
				ws.ReadJSON(&request)
				follow(friends, request.User, request.Followee)
				getFollowers(friends, 1)
				notifier <- request
			}
		}()
	*/

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

	for {
		//fmt.Println("running2")
		//var request FollowReq
		//ws.ReadJSON(&request)
		//follow(friends, request.User, request.Followee)
		//getFollowers(friends, 1)
		//notifier <- request
	}

	/*
		for {
			var request TweetPost
			ws.ReadJSON(&request)
			if request.Action == "post" {
				posts <- request
			}
		}
	*/
}
