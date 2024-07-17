package services

import (
	"back-end/data/request"
	"back-end/model"
	"back-end/repository"
	"back-end/sse"
	"back-end/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type FollowServiceImpl struct {
	FollowRepository repository.FollowRepository
	UserRepository   repository.UserRepository
	Validate         *validator.Validate
	h                *sse.NotificationSSE
}

func NewFollowServiceImpl(FollowRepository repository.FollowRepository, UserRepository repository.UserRepository, Validate *validator.Validate, h *sse.NotificationSSE) *FollowServiceImpl {
	return &FollowServiceImpl{FollowRepository: FollowRepository, UserRepository: UserRepository, Validate: Validate, h: h}

}

func (f FollowServiceImpl) Create(follow request.FollowRequest) error {
	err := f.Validate.Struct(follow)
	if err != nil {
		return err
	}

	follows := model.Follow{
		FollowerId:  follow.FollowerID,
		FollowingId: follow.FollowID,
	}

	err = f.FollowRepository.Create(follows)
	if err != nil {
		return err
	}

	user, err := f.UserRepository.FindUserByID(follow.FollowID)
	if err != nil {
		return err
	}

	if user.NotificationSetting.WebFollower == true {
		if _, exists := f.h.NotificationChannel[follow.FollowID]; !exists {
			f.h.NotificationChannel[follow.FollowID] = make(chan model.Notification)
		}
		f.h.NotificationChannel[follow.FollowID] <- model.Notification{
			NotifyId: utils.GenerateUUID(),
			UserId:   follow.FollowerID,
			Title:    "You have been followed",
			Body:     user.Username + " have Followed You at " + time.Now().Format("2006-01-02"),
			Status:   "OK",
			ReadAt:   time.Now(),
		}
	}

	fmt.Println(user.NotificationSetting.EmailFollower)
	fmt.Println(user.NotificationSettingId)
	if user.NotificationSetting.EmailFollower == true {
		err = utils.SendEmail(user.Email, "You have been followed", user.Username+" have Followed You at "+time.Now().Format("2006-01-02"))
	}

	return err
}

func (f FollowServiceImpl) GetFollowing(followerID string) ([]model.Follow, error) {
	fmt.Println("fl " + followerID)
	res, err := f.FollowRepository.GetFollowing(followerID)
	return res, err
}

func (f FollowServiceImpl) GetFollower(followingID string) ([]model.Follow, error) {
	res, err := f.FollowRepository.GetFollower(followingID)
	return res, err
}

func (f FollowServiceImpl) DeleteFollow(follow request.FollowRequest) error {
	err := f.Validate.Struct(follow)
	if err != nil {
		return err
	}

	follows := model.Follow{
		FollowerId:  follow.FollowerID,
		FollowingId: follow.FollowID,
	}

	err = f.FollowRepository.DeleteFollow(follows)
	return err
}

func (f FollowServiceImpl) GetMutualFollowing(followerID string) ([]model.Follow, error) {
	var mutual []model.Follow

	res, err := f.FollowRepository.GetFollowing(followerID)
	if err != nil {
		return res, err
	}
	for _, follow := range res {
		result, err := f.FollowRepository.GetMutualFollowing(follow.FollowingId, followerID)
		if err != nil {
			continue
		}
		mutual = append(mutual, result...)
	}

	return mutual, err
}
