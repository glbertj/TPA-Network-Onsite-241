package request

import "time"

type UserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type UserUpdateRequest struct {
	UserId  string    `json:"userId" validate:"required"`
	Gender  string    `json:"gender" validate:"required"`
	Country string    `json:"country" validate:"required"`
	Dob     time.Time `json:"dob" validate:"required"`
}

type ResetPasswordRequest struct {
	UserId   string `json:"userId" validate:"required"`
	Password string `json:"password" validate:"required"`
}
