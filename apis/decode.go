package apis

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"projects.com/apps/twitter-app/data"
)

type RequestHandler func(conn *websocket.Conn, req *data.RequestDecode, user_id data.ClientID) int

var ActionHandlers = map[string]RequestHandler{
	data.FollowAction:          FollowRequest_v2,
	data.PostAction:            PostRequest_v2,
	data.PostsByFolloweeAction: PostsByFollowees_v2,
	data.Subscribe:             SubscribeRequest,
}

var (
	FollowNotifier = make(data.FollowshipNotifier)
	PostNotifier   = make(data.PostingNotifier)
)

func FollowRequest(conn *websocket.Conn, req *data.RequestDecode, user_id data.ClientID) int {

	// Add Friend
	if data.Friends[req.FollowRequestDetails.CurrentUserId] == nil {
		data.Friends[req.FollowRequestDetails.CurrentUserId] = make(map[int]bool)
	}
	data.Friends[req.FollowRequestDetails.CurrentUserId][req.FollowRequestDetails.Followee] = true

	// Send Response
	conn.WriteJSON(req.FollowRequestDetails)

	// Send To Queue
	FollowNotifier <- *req.FollowRequestDetails

	// Get Followers
	data.GetFollowers(data.Friends, req.FollowRequestDetails.CurrentUserId)

	return 0
}

func FollowRequest_v2(conn *websocket.Conn, req *data.RequestDecode, user_id data.ClientID) int {

	// Add Friend
	if data.FriendList[user_id] == nil {
		data.FriendList[user_id] = make(map[data.ClientID]bool)
	}
	data.FriendList[user_id][data.RegisteredUsers[data.UserName(req.FollowRequestDetails.FolloweeName)].User_Id] = true // second index should be user_id of followee which is mapped to userName of followe

	// Send Response
	conn.WriteJSON(req.FollowRequestDetails)

	// Send To Queue
	FollowNotifier <- *req.FollowRequestDetails

	// Get Followers
	data.GetFollowers_v2(user_id)

	return 0
}

func PostRequest(conn *websocket.Conn, req *data.RequestDecode, user_id data.ClientID) int {

	// Time
	data.TimeInstance--

	// Added a Post
	PostNotifier <- *req.PostRequestDetails

	// Added to Mem DB
	data.Posts[data.TimeInstance] = *req.PostRequestDetails

	// Send Response
	conn.WriteJSON(req.PostRequestDetails)

	// Show Posts
	data.GetPosts(data.Posts)

	return 0
}

func PostRequest_v2(conn *websocket.Conn, req *data.RequestDecode, user_id data.ClientID) int {

	// Added to Mem DB
	data.Posts[data.TimeInstance] = *req.PostRequestDetails
	postedContent := data.Post{User_Id: user_id, Time_Post: data.TimeStamp(time.Now()), Content: data.PostContent(req.PostRequestDetails.ContentPost)}
	data.PostsList = append(data.PostsList, postedContent)

	// Added a Post
	data.PostsNotifier <- postedContent

	// Send Response
	conn.WriteJSON(req.PostRequestDetails)

	// Show Posts
	data.GetPosts_v2()

	return 0
}

func PostsByFollowees(conn *websocket.Conn, req *data.RequestDecode, user_id data.ClientID) int {

	// Posts By Followees
	current_user := req.PostsByFolloweesDetails.CurrentUserId

	// Write Back Posts By Followees
	for timeIns := range data.Posts {

		if data.Friends[current_user][data.Posts[timeIns].CurrentUserId] == true {
			response := data.PostedNotification{Action: "posts_by_followee", FolloweeUserID: "444", ContentPost: data.PostContent("Hi")}
			conn.WriteJSON(response)
		}
	}

	return 0

}

func PostsByFollowees_v2(conn *websocket.Conn, req *data.RequestDecode, user_id data.ClientID) int {

	// Write Back Posts By Followees
	for _, post := range data.PostsList {

		if data.FriendList[user_id][post.User_Id] == true {
			response := data.PostedNotification{Action: "posts_by_followee", FolloweeUserID: post.User_Id, ContentPost: post.Content}
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
