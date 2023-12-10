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
	User_Id   ClientID
	Email_Id  EmailID
	User_Name UserName
}

// Session
type SessionID string

type ActiveUser struct {
	Websocket_Session *websocket.Conn
	User_Id           ClientID
	Session_Id        SessionID
}

// Post
type TimeStamp time.Time
type PostContent string

type Post struct {
	User_Id   ClientID
	Time_Post TimeStamp
	Content   PostContent
}

// Maps
var ActiveSessions = make(map[SessionID]ActiveUser)
var RegisteredUsers = make(map[UserName]SignedUser)
var FriendList = make(map[ClientID]map[ClientID]bool)
var PostsList []Post

// Channels
var SessionNotifier = make(chan ActiveUser)
var PostsNotifier = make(chan Post)

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

func GetFollowers_v2(user_id ClientID) {

	fmt.Println("For user ID: ", user_id)
	for follower := range FriendList[user_id] {
		fmt.Printf("... above user follows %s with status %t\n", follower, FriendList[user_id][follower])
	}
}

func GetPosts(PostList PostsMemory) {

	for timeIns := range PostList {
		fmt.Printf("Post: %s has been posted by %d at time %d\n", PostList[timeIns].ContentPost, PostList[timeIns].CurrentUserId, timeIns)
	}
}

func GetPosts_v2() {

	for _, post := range PostsList {
		fmt.Printf("Post: %s has been posted by %s at time %d\n", post.Content, post.User_Id, post.Time_Post)
	}
}
