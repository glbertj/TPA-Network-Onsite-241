package services

import (
	"back-end/data/request"
	"back-end/data/response"
)

type UserService interface {
	Authenticate(req request.AuthRequest) (res response.AuthResponse, err error)
	GetCurrentUser(token string) (res response.UserResponse, err error)
	Register(req request.UserRequest) (res response.RegisterResponse, err error)
	LoginWithGoogle(req request.GoogleRequest) (res response.AuthResponse, err error)
	UpdateVerificationStatus(id string) (err error)
	GetUserById(id string) (res response.UserResponse, err error)
	UpdateUserProfile(req request.UserUpdateRequest) (res response.UserResponse, err error)
	GetAllUser() (res []response.UserResponse, err error)
	ForgotPassword(email string) (res response.RegisterResponse, err error)
	GetUserByVerifyLink(id string) (res response.UserResponse, err error)
	ChangePassword(password string, userId string) error
	Logout(userId string) error
	UpdateProfilePicture(userId string, avatar string) error
}
