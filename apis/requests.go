package apis

const (
	FollowAction = "follow"
	PostAction   = "post"
)

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
