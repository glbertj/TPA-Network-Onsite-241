package repository

import "back-end/model"

type NotificationSettingRepository interface {
	Create(notificationSetting model.NotificationSetting) error
	GetSettingBySettingId(id string) (res model.NotificationSetting, err error)
	UpdateNotificationSetting(userID string, setting model.NotificationSetting) (err error)
}
