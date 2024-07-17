package response

type FollowResponse struct {
	FollowerId  string       `json:"follower_id"`
	FollowingId string       `json:"following_id"`
	Follower    UserResponse `json:"follower" gorm:"foreignKey:FollowerId;references:UserId"`
	Following   UserResponse `json:"following" gorm:"foreignKey:FollowingId;references:UserId"`
}
