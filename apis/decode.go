package apis

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"projects.com/apps/twitter-app/data"
)

type RequestHandler func(conn *websocket.Conn, req *data.RequestDecode) int

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

func FollowRequest(conn *websocket.Conn, req *data.RequestDecode) int {

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

func PostRequest(conn *websocket.Conn, req *data.RequestDecode) int {

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

func PostsByFollowees(conn *websocket.Conn, req *data.RequestDecode) int {

	// Posts By Followees
	current_user := req.PostsByFolloweesDetails.CurrentUserId

	// Write Back Posts By Followees
	for timeIns := range data.Posts {

		if data.Friends[current_user][data.Posts[timeIns].CurrentUserId] == true {
			response := data.PostedNotification{Action: "posts_by_followee", Followee: data.Posts[timeIns].CurrentUserId, ContentPost: data.Posts[timeIns].ContentPost}
			conn.WriteJSON(response)
		}
	}

	return 0

}

func SubscribeRequest(conn *websocket.Conn, req *data.RequestDecode) int {

	fmt.Println("Add subscribe logic...")

	response := data.SubscribeRequestParams{CurrentUserId: req.SubscribeRequestDetails.CurrentUserId, Subscription: req.SubscribeRequestDetails.Subscription}

	conn.WriteJSON(response)

	return 0
}

func Broadcast() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		defer wg.Done()
		for {
			followNotification := <-FollowNotifier
			fmt.Printf("NOTIFICATION001: User ID %d has followed %d\n", followNotification.CurrentUserId, followNotification.Followee)

			notification := data.FollowNotification{Action: "FollowFeed", Follower: followNotification.CurrentUserId, Followee: followNotification.Followee}

			for _, wsclients := range data.Clients {
				wsclients.WriteJSON(notification)
			}
		}

	}()

	wg.Add(1)

	go func() {

		defer wg.Done()
		for {
			postNotification := <-PostNotifier
			fmt.Printf("NOTIFICATION002: User ID %d has posted %s\n", postNotification.CurrentUserId, postNotification.ContentPost)

			notification := data.PostedNotification{Action: "PostFeed", Followee: postNotification.CurrentUserId, ContentPost: postNotification.ContentPost}

			for _, wsclients := range data.Clients {
				wsclients.WriteJSON(notification)
			}
		}

	}()

	go func() {

		defer wg.Done()
		for {

			sessionAdded := <-data.SessionNotifier
			fmt.Printf("NOTIFICATION003: New Session added with session-id %s and user-id %s\n", sessionAdded.Session_Id, sessionAdded.User_Id)

			for _, sessionDetails := range data.ActiveSessions {
				fmt.Printf("NOTIFICATION003: Existing Session added with session-id %s and user-id %s\n", sessionDetails.Session_Id, sessionDetails.User_Id)
			}

		}

	}()

	wg.Wait()
}
