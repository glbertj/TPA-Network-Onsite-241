package repository

import (
	"back-end/model"
)

type FollowRepository interface {
	Create(follow model.Follow) error
	GetFollowing(followerID string) ([]model.Follow, error)
	GetFollower(followingID string) ([]model.Follow, error)
	DeleteFollow(follow model.Follow) error
	GetMutualFollowing(followerID string, followingID string) ([]model.Follow, error)
}
