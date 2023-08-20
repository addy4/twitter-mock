package apis

import (
	"github.com/gorilla/websocket"
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

func FollowRequest(conn *websocket.Conn, req *RequestDecode) int {
	//data.Friends[req.FollowRequestDetails.CurrentUserId][req.FollowRequestDetails.ToBeFollowedUserId] = true
	conn.WriteJSON(req.FollowRequestDetails)
	return 0
}

func PostRequest(conn *websocket.Conn, req *RequestDecode) int {
	return 0
}
