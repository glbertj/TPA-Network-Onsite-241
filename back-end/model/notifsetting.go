package model

type NotificationSetting struct {
	NotificationSettingId string `gorm:"primaryKey"`
	EmailFollower         bool
	EmailAlbum            bool
	WebFollower           bool
	WebAlbum              bool
}
