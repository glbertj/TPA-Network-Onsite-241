package response

type RegisterResponse struct {
	UserId     string `json:"userId"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	VerifyLink string `json:"verifyLink"`
}
