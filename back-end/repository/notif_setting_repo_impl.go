package repository

import (
	"back-end/database"
	"back-end/model"
	"back-end/utils"
	"fmt"
	"gorm.io/gorm"
)

type NotificationSettingRepositoryImpl struct {
	DB  *gorm.DB
	rdb *database.Redis
}

func NewNotificationSettingRepositoryImpl(DB *gorm.DB, rdb *database.Redis) *NotificationSettingRepositoryImpl {
	return &NotificationSettingRepositoryImpl{DB: DB, rdb: rdb}
}

func (n NotificationSettingRepositoryImpl) Create(notificationSetting model.NotificationSetting) error {
	err := n.DB.Create(&notificationSetting).Error
	return err
}

func (n NotificationSettingRepositoryImpl) GetSettingBySettingId(id string) (res model.NotificationSetting, err error) {
	err = n.DB.Where("notification_setting_id = ?", id).Find(&res).Error
	return
}

func (n NotificationSettingRepositoryImpl) UpdateNotificationSetting(userID string, setting model.NotificationSetting) (err error) {
	err = n.rdb.Del(utils.CurrentUserKey + userID)
	if err != nil {
		return err
	}
	fmt.Println(setting.NotificationSettingId)
	err = n.DB.Save(&setting).Error
	return
}
