package repository

import (
	"back-end/model"
	"time"
)

type UserRepository interface {
	Save(user model.User) error
	FindAll() ([]model.User, error)
	FindUserByID(id string) (user model.User, err error)
	FindByEmail(email string) (user model.User, err error)
	FindByEmailAndVerified(email string, verified bool) (user model.User, err error)
	Update(user model.User) (err error)
	UpdateRole(userId string) (err error)
	Delete(user model.User)
	GetCurrentUser(userId string) (user model.User, err error)
	UpdateGoogleId(userId string, email string, googleId string) (err error)
	GetUserByVerifyLink(verifyLink string) (user model.User, err error)
	UpdateVerifyLink(userId string, verifyLink string) (err error)
	UpdateRegister(userId string, verifyLink string, username string, hash []byte) (err error)
	ChangePassword(password []byte, userId string) error
	Logout(userId string) error
	UpdateProfilePicture(userId string, avatar string) error
	UpdateProfile(userId string, dob time.Time, country string, gender string) error
}
