package data

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Initialization
var TimeInstance = 100
var Flag = 0

// Data Containers
type PostsMemory map[int]PostRequestParams
type FriendsMemory map[int]map[int]bool
type ClientList []*websocket.Conn

var (
	Friends = make(FriendsMemory)
	Posts   = make(PostsMemory)
	Clients ClientList
)

// Data Methods
func GetFollowers(friendList FriendsMemory, userID int) {

	fmt.Println("For user ID: ", userID)
	for follower := range friendList[userID] {
		fmt.Printf("... above user follows %d with status\n", follower)
	}
}

func GetPosts(PostList PostsMemory) {

	for timeIns := range PostList {
		fmt.Printf("Post: %s has been posted by %d at time %d\n", PostList[timeIns].ContentPost, PostList[timeIns].CurrentUserId, timeIns)
	}
}
