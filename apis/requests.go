package apis

const (
	FollowAction = "follow"
	PostAction   = "post"
)

type FollowshipNotifier chan FollowRequestParams

type CommonAPI struct {
	Action string `json:"action"`
}

type FollowRequestParams struct {
	CurrentUserId      int `json:"current_user"`
	ToBeFollowedUserId int `json:"followed_user"`
}

type PostRequestParams struct {
	CurrentUserId      int `json:"current_user"`
	ToBeFollowedUserId int `json:"followed_user"`
}

type FollowNotification struct {
	Action         int `json:"action"`
	FollowedByUser int `json:"followed_by_user"`
	FollowedUser   int `json:"followed_user"`
}
