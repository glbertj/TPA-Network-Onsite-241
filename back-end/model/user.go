package model

import (
	"back-end/data/response"
	"time"
)

type User struct {
	UserId                string                               `gorm:"primaryKey"`
	Username              string                               `gorm:"not null"`
	Password              *[]byte                              `gorm:""`
	GoogleId              *string                              `gorm:""`
	Role                  string                               `gorm:"not null"`
	VerifiedAt            *time.Time                           `gorm:""`
	Email                 string                               `gorm:"not null"`
	Gender                *string                              `gorm:""`
	Country               *string                              `gorm:""`
	Avatar                *string                              `gorm:""`
	Dob                   *time.Time                           `gorm:""`
	VerifyLink            *string                              `gorm:""`
	NotificationSettingId string                               `gorm:""`
	NotificationSetting   response.NotificationSettingResponse `gorm:"foreignKey:NotificationSettingId;references:NotificationSettingId"`
}
