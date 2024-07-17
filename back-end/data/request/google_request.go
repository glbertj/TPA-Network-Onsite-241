package request

type GoogleRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	GoogleId string `json:"googleId"`
}
