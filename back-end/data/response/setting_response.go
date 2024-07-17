package response

type NotificationSettingResponse struct {
	NotificationSettingId string `json:"notificationSettingId"`
	EmailFollower         bool   `json:"emailFollower"`
	EmailAlbum            bool   `json:"emailAlbum"`
	WebFollower           bool   `json:"webFollower"`
	WebAlbum              bool   `json:"webAlbum"`
}

func (NotificationSettingResponse) TableName() string {
	return "notification_settings"
}
