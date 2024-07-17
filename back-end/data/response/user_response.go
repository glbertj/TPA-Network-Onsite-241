package response

import "time"

type UserResponse struct {
	UserId                string                      `json:"user_id"`
	Username              string                      `json:"username"`
	Email                 string                      `json:"email" `
	Role                  string                      `json:"role"`
	Avatar                *string                     `json:"avatar"`
	Country               *string                     `json:"country"`
	Gender                *string                     `json:"gender"`
	Dob                   *time.Time                  `json:"dob"`
	NotificationSettingId string                      `json:"notificationSettingId"`
	NotificationSetting   NotificationSettingResponse `json:"notificationSetting" gorm:"foreignKey:NotificationSettingId;references:NotificationSettingId"`
}

func (UserResponse) TableName() string {
	return "users"
}
