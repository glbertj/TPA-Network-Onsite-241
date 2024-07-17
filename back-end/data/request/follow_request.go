package request

type FollowRequest struct {
	FollowerID string `json:"followerId" validate:"required"`
	FollowID   string `json:"followId" validate:"required"`
}
