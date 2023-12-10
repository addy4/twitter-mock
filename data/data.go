package data

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

/* ---------- User Log In & Registeration ---------- */

type ClientID string
type EmailID string
type UserName string

type SignedUser struct {
	User_Id   ClientID
	Email_Id  EmailID
	User_Name UserName
}

/* ---------- Session Handling ---------- */

type SessionID string

type ActiveUser struct {
	Websocket_Session *websocket.Conn
	User_Id           ClientID
	Session_Id        SessionID
}

/* ---------- Post Content ---------- */

type TimeStamp time.Time
type PostContent string

type Post struct {
	User_Id   ClientID
	Time_Post TimeStamp
	Content   PostContent
}

/* ---------- Storage --------- */

var ActiveSessions = make(map[SessionID]ActiveUser)
var RegisteredUsers = make(map[UserName]SignedUser)
var FriendList = make(map[ClientID]map[ClientID]bool)
var PostsList []Post

type PostsMemory map[int]PostRequestParams
type FriendsMemory map[int]map[int]bool
type ClientList []*websocket.Conn

/* ---------- Notify & Communicate ---------- */

var SessionNotifier = make(chan ActiveUser)
var PostsNotifier = make(chan PostedNotification)
var FollowNotifier = make(chan FollowNotification)

/* ---------- Methods ---------- */
func GetFollowers(user_id ClientID) {

	fmt.Println("For user ID: ", user_id)
	for follower := range FriendList[user_id] {
		fmt.Printf("... above user follows %s with status %t\n", follower, FriendList[user_id][follower])
	}
}

func GetPosts() {

	for _, post := range PostsList {
		fmt.Printf("Post: %s has been posted by %s at time %d\n", post.Content, post.User_Id, post.Time_Post)
	}
}
