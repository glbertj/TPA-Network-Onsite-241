package repository

import (
	"back-end/database"
	"back-end/model"
	"fmt"
	"gorm.io/gorm"
)

type FollowRepositoryImpl struct {
	DB  *gorm.DB
	rdb *database.Redis
}

func NewFollowRepositoryImpl(DB *gorm.DB, rdb *database.Redis) *FollowRepositoryImpl {
	return &FollowRepositoryImpl{DB: DB, rdb: rdb}
}

func (f FollowRepositoryImpl) Create(follow model.Follow) error {
	err := f.DB.Create(&follow).Error
	return err
}

func (f FollowRepositoryImpl) GetFollowing(followerID string) (res []model.Follow, err error) {
	fmt.Println(followerID)
	err = f.DB.Where("follower_id = ?", followerID).Preload("Follower").Preload("Follower.NotificationSetting").Preload("Following").Preload("Following.NotificationSetting").Find(&res).Error
	return
}

func (f FollowRepositoryImpl) GetFollower(followingID string) (res []model.Follow, err error) {
	err = f.DB.Where("following_id = ?", followingID).Preload("Follower").Preload("Follower.NotificationSetting").Preload("Following").Preload("Following.NotificationSetting").Find(&res).Error
	return
}

func (f FollowRepositoryImpl) DeleteFollow(follow model.Follow) error {
	err := f.DB.Delete(&follow).Error
	return err
}

func (f FollowRepositoryImpl) GetMutualFollowing(followerID string, followingID string) (res []model.Follow, err error) {
	err = f.DB.Where("follower_id = ? AND following_id = ?", followerID, followingID).Preload("Follower").Preload("Following").Find(&res).Error
	return
}
