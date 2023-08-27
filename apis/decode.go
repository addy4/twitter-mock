package apis

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"projects.com/apps/twitter-app/data"
)

type RequestDecode struct {
	FollowRequestDetails *FollowRequestParams `json:"follow,omitempty"`
	PostRequestDetails   *PostRequestParams   `json:"post,omitempty"`
}

type RequestHandler func(conn *websocket.Conn, req *RequestDecode) int

var ActionHandlers = map[string]RequestHandler{
	FollowAction: FollowRequest,
	PostAction:   PostRequest,
}

var (
	FollowNotifier = make(FollowshipNotifier)
	PostNotifier   = make(PostingNotifier)
)

func FollowRequest(conn *websocket.Conn, req *RequestDecode) int {

	// Add Friend
	if data.Friends[req.FollowRequestDetails.CurrentUserId] == nil {
		data.Friends[req.FollowRequestDetails.CurrentUserId] = make(map[int]bool)
	}
	data.Friends[req.FollowRequestDetails.CurrentUserId][req.FollowRequestDetails.ToBeFollowedUserId] = true

	// Send Response
	conn.WriteJSON(req.FollowRequestDetails)

	// Send To Queue
	FollowNotifier <- *req.FollowRequestDetails

	// Notify
	data.GetFollowers(data.Friends, req.FollowRequestDetails.CurrentUserId)

	return 0
}

func PostRequest(conn *websocket.Conn, req *RequestDecode) int {

	// Added a Post
	PostNotifier <- *req.PostRequestDetails

	// Send Response
	conn.WriteJSON(req.PostRequestDetails)

	return 0
}

func Broadcast() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		defer wg.Done()
		for {
			followNotification := <-FollowNotifier
			fmt.Printf("NOTIFICATION: User ID %d has followed %d", followNotification.CurrentUserId, followNotification.ToBeFollowedUserId)

			notification := FollowNotification{Action: "FollowFeed", FollowedByUser: followNotification.CurrentUserId, FollowedUser: followNotification.ToBeFollowedUserId}

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
			fmt.Printf("NOTIFICATION: User ID %d has posted %s", postNotification.CurrentUserId, postNotification.ContentPost)

			notification := PostedNotification{Action: "PostFeed", Follower: postNotification.CurrentUserId, ContentData: postNotification.ContentPost}

			for _, wsclients := range data.Clients {
				wsclients.WriteJSON(notification)
			}
		}

	}()

	wg.Wait()
}
