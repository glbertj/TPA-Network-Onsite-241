package request

type PlayListRequest struct {
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
