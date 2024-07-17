package services

import (
	"back-end/data/request"
	"back-end/model"
)

type NotificationSettingService interface {
	CreateNotificationSetting(userId string) error
	GetSettingBySettingId(id string) (res model.NotificationSetting, err error)
	UpdateNotificationSetting(setting request.NotificationSettingRequest) (err error)
}
