package data

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type FriendsMemory map[int]map[int]bool
type ClientList []*websocket.Conn

var (
	Friends = make(FriendsMemory)
	Clients ClientList
)

func GetFollowers(friendList FriendsMemory, userID int) {

	fmt.Println("For user ID: ", userID)
	for follower := range friendList[userID] {
		fmt.Printf("... above user follows %d with status\n", follower)
	}
}
