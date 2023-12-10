package apis

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"projects.com/apps/twitter-app/data"
)

type RequestHandler func(conn *websocket.Conn, req *data.RequestDecode, user_id data.ClientID) int

var ActionHandlers = map[string]RequestHandler{
	data.FollowAction:          FollowRequest,
	data.PostAction:            PostRequest,
	data.PostsByFolloweeAction: PostsByFollowees,
	data.Subscribe:             SubscribeRequest,
}

var (
	FollowNotifier = make(data.FollowshipNotifier)
	PostNotifier   = make(data.PostingNotifier)
)

func FollowRequest(conn *websocket.Conn, req *data.RequestDecode, user_id data.ClientID) int {

	// Add Friend
	if data.FriendList[user_id] == nil {
		data.FriendList[user_id] = make(map[data.ClientID]bool)
	}
	data.FriendList[user_id][data.RegisteredUsers[data.UserName(req.FollowRequestDetails.FolloweeName)].User_Id] = true // second index should be user_id of followee which is mapped to userName of followee

	// For Notification...
	followNotification := data.FollowNotification{Action: data.FollowFeed, CurrentUser: user_id, Followee: req.FollowRequestDetails.FolloweeName}
	data.FollowNotifier <- followNotification

	// Send Response
	conn.WriteJSON(req.FollowRequestDetails)

	return 0
}

func PostRequest(conn *websocket.Conn, req *data.RequestDecode, user_id data.ClientID) int {

	// Added to Mem DB
	postedContent := data.Post{User_Id: user_id, Time_Post: data.TimeStamp(time.Now()), Content: data.PostContent(req.PostRequestDetails.ContentPost)}
	data.PostsList = append(data.PostsList, postedContent)

	// For Notification...
	postNotification := data.PostedNotification{Action: data.PostFeed, CurrentUser: user_id, ContentPost: data.PostContent(req.PostRequestDetails.ContentPost)}
	data.PostsNotifier <- postNotification

	// Send Response
	conn.WriteJSON(req.PostRequestDetails)

	return 0
}

func PostsByFollowees(conn *websocket.Conn, req *data.RequestDecode, user_id data.ClientID) int {

	// Write Back Posts By Followees
	for _, post := range data.PostsList {

		if data.FriendList[user_id][post.User_Id] == true {
			response := data.PostedNotification{Action: "posts_by_followee", CurrentUser: post.User_Id, ContentPost: post.Content}
			conn.WriteJSON(response)
		}
	}

	return 0

}

func SubscribeRequest(conn *websocket.Conn, req *data.RequestDecode, user_id data.ClientID) int {

	fmt.Println("Add subscribe logic...")

	response := data.SubscribeRequestParams{CurrentUserId: req.SubscribeRequestDetails.CurrentUserId, Subscription: req.SubscribeRequestDetails.Subscription}

	conn.WriteJSON(response)

	return 0
}
