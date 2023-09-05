package data

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

// Initialization
var TimeInstance = 100
var Flag = 0

// User
type ClientID string
type EmailID string
type UserName string

type SignedUser struct {
	user_id   ClientID
	email_id  EmailID
	user_name UserName
}

// Session
type SessionID string
type WebsocketSession *websocket.Conn

type ActiveUser struct {
	websocket_session WebsocketSession
	user_id           ClientID
}

// Post
type TimeStamp time.Time

type Post struct {
	user_id    ClientID
	time_stamp TimeStamp
	content    string
}

// Maps
var ActiveSessions = make(map[SessionID]ActiveUser)
var RegisteredUsers = make(map[UserName]SignedUser)
var FriendList = make(map[ClientID]map[ClientID]bool)
var PostsList []Post

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
