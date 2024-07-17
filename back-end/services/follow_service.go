package services

import (
	"back-end/data/request"
	"back-end/model"
)

type FollowService interface {
	Create(follow request.FollowRequest) error
	GetFollowing(followerID string) ([]model.Follow, error)
	GetFollower(followingID string) ([]model.Follow, error)
	DeleteFollow(follow request.FollowRequest) error
	GetMutualFollowing(followerID string) ([]model.Follow, error)
}
