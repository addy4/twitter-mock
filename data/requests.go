package data

const (
	FollowAction          = "follow"
	PostAction            = "post"
	PostsByFolloweeAction = "posts_by_followees"
	Subscribe             = "subscribe"
)

type FollowshipNotifier chan FollowRequestParams
type PostingNotifier chan PostRequestParams

type RequestDecode struct {
	FollowRequestDetails    *FollowRequestParams    `json:"follow,omitempty"`
	PostRequestDetails      *PostRequestParams      `json:"post,omitempty"`
	PostsByFolloweesDetails *PostsByFolloweesParams `json:"posts_by_followees"`
	SubscribeRequestDetails *SubscribeRequestParams `json:"subscribe"`
}

type CommonAPI struct {
	Action string `json:"action"`
}

type FollowRequestParams struct {
	CurrentUserId int    `json:"current_user"`
	Followee      int    `json:"followee"`
	FolloweeName  string `json:"followeeName"`
}

type PostRequestParams struct {
	CurrentUserId int    `json:"current_user"`
	ContentPost   string `json:"content"`
}

type SubscribeRequestParams struct {
	CurrentUserId int    `json:"current_user"`
	Subscription  string `json:"subscription"`
}

type FollowNotification struct {
	Action   string `json:"action"`   // follow_feed
	Follower int    `json:"follower"` // "X" followed Y
	Followee int    `json:"followee"` // "Y" has been followed by X
}

type PostedNotification struct {
	Action         string      `json:"action"`           // post_feed
	FolloweeUserID ClientID    `json:"followee_user_id"` // "X" posted Y
	ContentPost    PostContent `json:"content"`          // "Y" has been posted by X
}

type PostsByFolloweesParams struct {
	CurrentUserId int `json:"current_user"`
}
