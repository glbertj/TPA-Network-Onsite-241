package model

import "back-end/data/response"

type Follow struct {
	FollowerId  string                `gorm:"primaryKey"`
	FollowingId string                `gorm:"primaryKey"`
	Follower    response.UserResponse `gorm:"foreignKey:FollowerId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Following   response.UserResponse `gorm:"foreignKey:FollowingId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
