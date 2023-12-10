package data

const (
	FollowAction          = "follow"
	PostAction            = "post"
	PostsByFolloweeAction = "posts_by_followees"
	Subscribe             = "subscribe"
	PostFeed              = "postFeed"
	FollowFeed            = "followFeed"
)

type FollowshipNotifier chan FollowRequestParams
type PostingNotifier chan PostRequestParams

// Request Decode
type RequestDecode struct {
	FollowRequestDetails    *FollowRequestParams    `json:"follow,omitempty"`
	PostRequestDetails      *PostRequestParams      `json:"post,omitempty"`
	PostsByFolloweesDetails *PostsByFolloweesParams `json:"posts_by_followees"`
	SubscribeRequestDetails *SubscribeRequestParams `json:"subscribe"`
}

// Action Decode
type CommonAPI struct {
	Action string `json:"action"`
}

/* ---------- API Request Messages ---------- */

// Follow Request
type FollowRequestParams struct {
	FolloweeName string `json:"followeeName"`
}

// Post Request
type PostRequestParams struct {
	ContentPost string `json:"content"`
}

// Subscribe Request
type SubscribeRequestParams struct {
	CurrentUserId int    `json:"current_user"`
	Subscription  string `json:"subscription"`
}

// Posts By Followees
type PostsByFolloweesParams struct {
	Foo bool `json:"foo"`
}

/* ---------- Notifications ---------- */

type FollowNotification struct {
	Action      string   `json:"action"`    // follow_feed
	CurrentUser ClientID `json:"client_id"` // "X" followed Y
	Followee    string   `json:"followee"`  // "Y" has been followed by X
}

type PostedNotification struct {
	Action      string      `json:"action"`    // post_feed
	CurrentUser ClientID    `json:"client_id"` // "X" posted Y
	ContentPost PostContent `json:"content"`   // "Y" has been posted by X
}
