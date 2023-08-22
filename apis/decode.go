package apis

import (
	"fmt"

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
	return 0
}

func Broadcast() {

	go func() {

		for {
			followNotification := <-FollowNotifier
			fmt.Printf("NOTIFICATION: User ID %d has followed %d", followNotification.CurrentUserId, followNotification.ToBeFollowedUserId)
		}

	}()

}
