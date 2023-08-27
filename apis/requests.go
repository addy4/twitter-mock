package apis

const (
	FollowAction = "follow"
	PostAction   = "post"
)

type FollowshipNotifier chan FollowRequestParams
type PostingNotifier chan PostRequestParams

type CommonAPI struct {
	Action string `json:"action"`
}

type FollowRequestParams struct {
	CurrentUserId      int `json:"current_user"`
	ToBeFollowedUserId int `json:"followed_user"`
}

type PostRequestParams struct {
	CurrentUserId int    `json:"current_user"`
	ContentPost   string `json:"content"`
}

type FollowNotification struct {
	Action         string `json:"action"`   // follow_feed
	FollowedByUser int    `json:"follower"` // "X" followed Y
	FollowedUser   int    `json:"followee"` // "Y" has been followed by X
}

type PostedNotification struct {
	Action      string `json:"action"`   // follow_feed
	Follower    int    `json:"follower"` // "X" followed Y
	ContentData string `json:"content"`  // "Y" has been followed by X
}
