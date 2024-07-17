package services

import (
	"back-end/data/request"
	"back-end/model"
	"back-end/repository"
	"back-end/utils"
	"github.com/go-playground/validator/v10"
)

type NotificationSettingServiceImpl struct {
	NotificationSettingRepository repository.NotificationSettingRepository
	Validate                      *validator.Validate
}

func NewNotificationSettingServiceImpl(NotificationSettingRepository repository.NotificationSettingRepository, Validate *validator.Validate) *NotificationSettingServiceImpl {
	return &NotificationSettingServiceImpl{NotificationSettingRepository: NotificationSettingRepository, Validate: Validate}
}

func (n *NotificationSettingServiceImpl) CreateNotificationSetting(userId string) error {
	notificationSetting := model.NotificationSetting{
		NotificationSettingId: utils.GenerateUUID(),
		EmailFollower:         false,
		EmailAlbum:            false,
		WebFollower:           false,
		WebAlbum:              false,
	}
	return n.NotificationSettingRepository.Create(notificationSetting)
}

func (n *NotificationSettingServiceImpl) GetSettingBySettingId(id string) (res model.NotificationSetting, err error) {
	return n.NotificationSettingRepository.GetSettingBySettingId(id)
}

func (n *NotificationSettingServiceImpl) UpdateNotificationSetting(setting request.NotificationSettingRequest) (err error) {
	err = n.Validate.Struct(setting)
	if err != nil {
		return err
	}
	settingModel := model.NotificationSetting{
		NotificationSettingId: setting.NotificationSettingId,
		EmailFollower:         setting.EmailFollower,
		EmailAlbum:            setting.EmailAlbum,
		WebFollower:           setting.WebFollower,
		WebAlbum:              setting.WebAlbum,
	}
	err = n.NotificationSettingRepository.UpdateNotificationSetting(setting.UserID, settingModel)
	if err != nil {
		return err
	}
	return nil

}
