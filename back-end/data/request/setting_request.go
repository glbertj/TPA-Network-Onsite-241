package request

type NotificationSettingRequest struct {
	UserID                string `json:"userId"`
	NotificationSettingId string `json:"notificationSettingId"`
	EmailFollower         bool   `json:"emailFollower"`
	EmailAlbum            bool   `json:"emailAlbum"`
	WebFollower           bool   `json:"webFollower"`
	WebAlbum              bool   `json:"webAlbum"`
}
